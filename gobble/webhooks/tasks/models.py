from typing import Literal

from pydantic import BaseModel, Field

from gobble.webhooks.event_types import EventType


class BaseTask(BaseModel):
    name: str
    event_types: list[EventType] | Literal["all"] = Field(
        "all", description="On which event types to call the webhook"
    )
