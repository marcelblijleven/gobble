import logging

from gobble.webhooks.event_types import EventType
from gobble.webhooks.tasks.registry import register_webhook_task
from gobble.webhooks.webhook import WebhookEvent

logger = logging.getLogger(__name__)


@register_webhook_task(*[type_ for type_ in EventType])
async def log_event(event: WebhookEvent) -> None:
    """
    Logs an incoming event to stdout

    Args:
        event: the webook event

    Returns:
        Nothing
    """
    logger.info(f"Received webhook: {event.event}")
