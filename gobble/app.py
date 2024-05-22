import logging
from contextlib import asynccontextmanager

from fastapi import FastAPI

from gobble.lib.plexapi.client import plex_client
from gobble.routes.v1.plex.router import plex_router


@asynccontextmanager
async def lifespan(app_: FastAPI):
    logging.basicConfig(level=logging.INFO)
    plex_server_identity = await plex_client.get_server_identity()
    plex_server_capabilities = await plex_client.get_server_capabilities()
    print(plex_server_capabilities, plex_server_identity)
    yield
    ...

app = FastAPI(lifespan=lifespan)
app.include_router(plex_router)
