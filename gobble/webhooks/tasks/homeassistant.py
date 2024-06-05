import logging
from typing import Any, Literal

from pydantic import AnyUrl, Field

from gobble.homeassistant.client import get_homeassistant_client
from gobble.webhooks.event_types import EventType
from gobble.webhooks.tasks.models import BaseTask, ComplexTask
from gobble.webhooks.tasks.registry import register_webhook_task
from gobble.webhooks.webhook import WebhookEvent

STATE: Literal["state"] = "state"
SERVICE: Literal["service"] = "service"

logger = logging.getLogger(__name__)


class HomeassistantServiceTask(BaseTask):
    entity_id: str
    domain: str
    service: str
    attributes: dict[str, Any] | None = Field(None)


class HomeassistantSettings(ComplexTask):
    access_token: str
    url: AnyUrl
    tasks: list[HomeassistantServiceTask] = Field(default_factory=list)  # type: ignore


@register_webhook_task(*[type_ for type_ in EventType])
async def update_homeassistant_entity(event: WebhookEvent) -> None:
    from gobble.config import settings

    if (ha := settings.tasks.homeassistant) is None:
        return

    client = get_homeassistant_client(ha.access_token, ha.url)

    for task in ha.tasks:
        if (event_types := task.event_types) == "all" or event.event in event_types:
            if isinstance(task, HomeassistantServiceTask):
                resp = await client.call_service(
                    task.entity_id, task.domain, task.service
                )
            else:
                logger.error("Unknown Homeassistant task received")
                continue

            logger.debug(f"Updated Homeassistant entity: {resp}")
