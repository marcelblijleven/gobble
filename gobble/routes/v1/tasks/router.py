import logging
from typing import Mapping

from fastapi import APIRouter

from gobble.config import settings
from gobble.routes.v1.tasks import models
from gobble.webhooks.tasks.models import BaseTask

logger = logging.getLogger(__name__)
tasks_router = APIRouter(prefix="/tasks", tags=["Tasks"])


@tasks_router.get("/", response_model=models.TasksResponse)
async def get_plex_version():
    tasks = settings.tasks
    tasks_dict: Mapping[str, str | BaseTask | list[BaseTask]] = {}

    for task_name in tasks.model_fields_set:
        task = getattr(tasks, task_name)

        if not task:
            value = "disabled"
        else:
            if isinstance(task, bool):
                value = "all"
            else:
                value = task

        tasks_dict[task_name] = value

    return models.TasksResponse(**tasks_dict)
