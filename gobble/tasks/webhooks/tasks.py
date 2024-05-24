import asyncio
import logging

from gobble.routes.v1.plex.models import WebhookEventModel
from gobble.tasks.webhooks.event_types import EventType
from gobble.tasks.webhooks.registry import webhook_task_registry

logger = logging.getLogger(__name__)


async def call_registered_webhook_tasks_for_event(
    event: WebhookEventModel, event_type: EventType
) -> None:
    """
    Will create a task group for the registered callbacks that belong to the provided event type.

    Args:
        event: a webhook event model
        event_type: the type of webhook event

    Returns:

    """
    async with asyncio.TaskGroup() as task_group:
        for task in webhook_task_registry.get(event_type, []):
            task_group.create_task(task(event))

    logger.debug(f"finished tasks for {event_type=}")
