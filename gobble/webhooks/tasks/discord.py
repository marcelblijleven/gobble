import asyncio
import logging
from enum import IntEnum

from pydantic import AnyHttpUrl, AnyUrl, Field

from gobble.discord.client import call_webhook
from gobble.discord.models import (
    DiscordWebhookBody,
    EmbedField,
    EmbedObject,
)
from gobble.webhooks.event_types import EventType
from gobble.webhooks.tasks.models import BaseTask
from gobble.webhooks.tasks.registry import register_webhook_task
from gobble.webhooks.webhook import MediaType, WebhookEvent

logger = logging.getLogger(__name__)


class DiscordSettings(BaseTask):
    webhook_url: AnyHttpUrl = Field(...)
    username: str | None = Field(
        "Gobble", description="The username to use in the webhook"
    )
    avatar_url: AnyUrl | None = Field(None)


class Colors(IntEnum):
    Play = 4717623
    Resume = 4717623
    Pause = 16566839
    Stop = 16529207
    Scrobble = 3648252
    Rate = 3669178


def _pad_index(index: int | None) -> str:
    """
    Pad a number with zero if less than 10

    Args:
        index: the integer to pad

    Returns:
        A padded string
    """
    if index is None:
        return ""

    return f"{index:02}"


def _get_event_action(source: str, event_type: EventType) -> str:
    return f"{source.capitalize()}: {" ".join(event_type.value.split("_"))}"


def _get_event_color(event_type: EventType) -> int:
    match event_type:
        case EventType.MediaPlay:
            return Colors.Play
        case EventType.MediaResume:
            return Colors.Resume
        case EventType.MediaPause:
            return Colors.Pause
        case EventType.MediaStop:
            return Colors.Stop
        case EventType.MediaScrobble:
            return Colors.Scrobble
        case EventType.MediaRate:
            return Colors.Rate
    return 0


def create_music_body(event: WebhookEvent):
    """
    Create a music specific discord webhook body
    Args:
        event: the incoming webhook event

    Returns:
        DiscordWebhookBody
    """
    return DiscordWebhookBody(
        content="",
        embeds=[
            EmbedObject(
                title=event.title,
                description=_get_event_action("Plex", event.event),
                fields=[
                    EmbedField(name="Artist", value=event.parent_title, inline=True),
                    EmbedField(
                        name="Album", value=event.grandparent_title, inline=True
                    ),
                    EmbedField(name="Username", value=event.username, inline=False),
                ],
            ),
        ],
    )


def create_show_body(event: WebhookEvent):
    """
    Create a tv show specific discord webhook body
    Args:
        event: the incoming webhook event

    Returns:
        DiscordWebhookBody
    """
    return DiscordWebhookBody(
        content="",
        embeds=[
            EmbedObject(
                title=event.grandparent_title,
                description=_get_event_action("Plex", event.event),
                fields=[
                    EmbedField(name="Name", value=event.title, inline=True),
                    EmbedField(
                        name="Season", value=_pad_index(event.parent_index), inline=True
                    ),
                    EmbedField(
                        name="Episode", value=_pad_index(event.index), inline=True
                    ),
                    EmbedField(name="Username", value=event.username, inline=False),
                ],
            ),
        ],
    )


def create_movie_body(event: WebhookEvent) -> DiscordWebhookBody:
    """
    Create a movie specific discord webhook body
    Args:
        event: the incoming webhook event

    Returns:
        DiscordWebhookBody
    """
    return DiscordWebhookBody(
        content="",
        embeds=[
            EmbedObject(
                title=event.title,
                description=_get_event_action("Plex", event.event),
                fields=[
                    EmbedField(name="Username", value=event.username, inline=False),
                ],
            ),
        ],
    )


@register_webhook_task(*[type_ for type_ in EventType])
async def call_discord_webhook(event: WebhookEvent) -> None:
    """
    Calls a Discord webhook

    Args:
        event: the webook event

    Returns:
        Nothing
    """
    from gobble.config import settings

    if not settings.tasks.discord:
        return

    async with asyncio.TaskGroup() as task_group:
        for discord_setting in settings.tasks.discord:
            if (
                event_types := discord_setting.event_types
            ) == "all" or event.event in event_types:
                if event.media_type == MediaType.Show:
                    body = create_show_body(event)
                elif event.media_type == MediaType.Movie:
                    body = create_movie_body(event)
                elif event.media_type == MediaType.Music:
                    body = create_music_body(event)
                else:
                    return

                task_group.create_task(
                    call_webhook(url=str(discord_setting.webhook_url), body=body)
                )
