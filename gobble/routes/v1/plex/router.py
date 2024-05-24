import logging
from functools import partial
from typing import Annotated

from fastapi import APIRouter, Form, BackgroundTasks, status
from pydantic import ValidationError

from gobble.exceptions import UnsupportedEventTypeException
from gobble.plexapi.client import plex_client
from gobble.routes.v1.plex import models
from gobble.routes.v1.plex.models import WebhookEventModel
from gobble.routes.v1.plex.utils import get_event_type
from gobble.tasks.webhooks.tasks import call_registered_webhook_tasks_for_event

logger = logging.getLogger(__name__)
plex_router = APIRouter(prefix="/plex", tags=["Plex"])


@plex_router.get("/version", response_model=models.VersionResponseModel)
async def get_plex_version():
    return {"version": plex_client.version}


@plex_router.post("/webhook", status_code=status.HTTP_200_OK)
async def webhook(payload: Annotated[bytes, Form()], background_tasks: BackgroundTasks):
    try:
        event = WebhookEventModel.model_validate_json(payload)
        event_type = get_event_type(event)
    except ValidationError as exc:
        logger.error(f"received unknown webhook data from Plex: {exc}")
        return
    except UnsupportedEventTypeException as exc:
        logger.error(f"{exc}")
        return

    task_for_event = partial(call_registered_webhook_tasks_for_event, event, event_type)
    background_tasks.add_task(task_for_event)
