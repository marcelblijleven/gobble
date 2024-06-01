from typing import Literal, Mapping, TypeAlias

from pydantic import RootModel

from gobble.webhooks.tasks.models import BaseTask

Disabled: TypeAlias = Literal["disabled"]


class TasksResponse(RootModel):
    root: Mapping[str, str | BaseTask | list[BaseTask]]
