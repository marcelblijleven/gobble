import logging.config
from contextlib import asynccontextmanager

from fastapi import FastAPI

from gobble.config import settings
from gobble.logging_config import LOGGING_CONFIG
from gobble.routes.v1.plex.router import plex_router
from gobble.utils import identify_servers

logging.config.dictConfig(LOGGING_CONFIG)


@asynccontextmanager
async def lifespan(app_: FastAPI):
    media_servers = await identify_servers(settings)

    yield {"media_servers": media_servers}

    ...


app = FastAPI(lifespan=lifespan)
app.include_router(plex_router)
