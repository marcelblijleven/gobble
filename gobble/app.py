import logging.config
from contextlib import asynccontextmanager

from fastapi import FastAPI

from gobble.config import settings
from gobble.logging_config import LOGGING_CONFIG
from gobble.routes.v1.plex.router import plex_router
from gobble.routes.v1.tasks.router import tasks_router
from gobble.utils import identify_servers
from gobble.webhooks.tasks.registry import autodetect

logging.basicConfig(level=logging.INFO)
logging.config.dictConfig(LOGGING_CONFIG)
logger = logging.getLogger(__name__)


@asynccontextmanager
async def lifespan(app_: FastAPI):
    autodetect()
    media_servers = await identify_servers(settings)

    yield {"media_servers": media_servers}

    ...


app = FastAPI(lifespan=lifespan)
app.include_router(plex_router)
app.include_router(tasks_router)
