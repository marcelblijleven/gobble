import logging
from functools import partial
from typing import Annotated

from fastapi import APIRouter, BackgroundTasks, Form, Request, status
from pydantic import ValidationError

from gobble.exceptions import UnsupportedEventTypeException
from gobble.plex.server import PlexServer
from gobble.protocols import MediaServer
from gobble.routes.v1.plex import models
from gobble.routes.v1.plex.models import PlexWebhookEventModel
from gobble.webhooks.tasks.registry import call_registered_webhook_tasks_for_event

logger = logging.getLogger(__name__)
plex_router = APIRouter(prefix="/plex", tags=["Plex"])


@plex_router.get("/version", response_model=models.VersionResponseModel)
async def get_plex_version(request: Request):
    media_servers: dict[str, MediaServer] = request.state.media_servers
    versions: dict[str, str] = {}

    for server in media_servers.values():
        if isinstance(server, PlexServer):
            versions[server.name] = server.version

    return versions


@plex_router.post("/webhook", status_code=status.HTTP_200_OK)
async def webhook(payload: Annotated[bytes, Form()], background_tasks: BackgroundTasks):
    try:
        event = PlexWebhookEventModel.model_validate_json(payload)
    except ValidationError as exc:
        logger.error(f"received unknown webhook data from Plex: {exc}")
        return
    except UnsupportedEventTypeException as exc:
        logger.error(f"{exc}")
        return

    webhook_event = event.to_generic_webhook_event()
    task_for_event = partial(call_registered_webhook_tasks_for_event, webhook_event)
    background_tasks.add_task(task_for_event)
