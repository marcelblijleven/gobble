import logging

from gobble.routes.v1.plex.models import WebhookEventModel
from gobble.tasks.webhooks.event_types import EventType
from gobble.tasks.webhooks.registry import register_webhook_task

logger = logging.getLogger(__name__)


@register_webhook_task(*[type_ for type_ in EventType])
async def log_event(event: WebhookEventModel) -> None:
    """
    Logs an incoming event to stdout

    Args:
        event: the webook event

    Returns:
        Nothing
    """
    logger.info(f"Received webhook: {event.event}")
