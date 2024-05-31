from typing import Literal

from pydantic import AnyHttpUrl, AnyUrl, BaseModel, Field

from gobble.webhooks.event_types import EventType


class DiscordSettings(BaseModel):
    webhook_url: AnyHttpUrl = Field(...)
    event_types: list[EventType] | Literal["all"] = Field(
        "all", description="On which event types to call the webhook"
    )
    username: str | None = Field(
        "Gobble", description="The username to use in the webhook"
    )
    avatar_url: AnyUrl | None = Field(None)
