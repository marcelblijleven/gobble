import logging
from typing import Annotated

from fastapi import APIRouter, Form
from pydantic import ValidationError

from gobble.lib.plexapi.client import plex_client
from gobble.routes.v1.plex import models
from gobble.routes.v1.plex.models import WebhookEventModel

logger = logging.getLogger(__name__)
plex_router = APIRouter(prefix="/plex", tags=["Plex"])


@plex_router.get("/version", response_model=models.VersionResponseModel)
async def get_plex_version():
    return {"version": plex_client.version}


@plex_router.post("/webhook")
async def webhook(payload: Annotated[bytes, Form()]):
    try:
        event = WebhookEventModel.model_validate_json(payload)
    except ValidationError as exc:
        logger.error(f"received unknown webhook data from Plex: {exc}")
        return

    logger.info(f"received webhook event from Plex: {event.event}")
