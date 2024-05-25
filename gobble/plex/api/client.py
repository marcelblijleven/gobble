from enum import Enum

import httpx
from httpx import Response

from gobble.plex.api.schemas import ServerCapabilitiesResponse, ServerIdentityResponse
from gobble.plex.config import PlexSettings


class HttpMethod(Enum):
    GET = "GET"
    POST = "POST"
    PATCH = "PATCH"


class PlexClient:
    def __init__(self, plex_settings: PlexSettings) -> None:
        self._settings = plex_settings
        self._default_headers = {
            "X-Plex-Token": self._settings.token,
            "Accept": "application/json,application/xml",
        }

    async def get_server_identity(self) -> ServerIdentityResponse:
        url = f"{self._settings.server_url}identity"
        resp = await self._make_async_request(url, method=HttpMethod.GET)

        return ServerIdentityResponse.model_validate(resp.json())

    async def get_server_capabilities(self) -> ServerCapabilitiesResponse:
        url = str(self._settings.server_url)
        resp = await self._make_async_request(url, method=HttpMethod.GET)

        return ServerCapabilitiesResponse.model_validate(resp.json())

    def get_server_identity_sync(self) -> ServerIdentityResponse:
        url = f"{self._settings.server_url}identity"
        resp = self._make_sync_request(url, method=HttpMethod.GET)

        server_identity = ServerIdentityResponse.model_validate(resp.json())
        self._server_identity = server_identity
        return self._server_identity

    async def _make_async_request(
        self, url: str, *, method: HttpMethod, **kwargs
    ) -> Response:
        async with httpx.AsyncClient() as client:
            headers = self._default_headers | kwargs.pop("headers", {})
            resp = await client.request(
                method=method.value, url=url, headers=headers, **kwargs
            )
            resp.raise_for_status()

        return resp

    def _make_sync_request(self, url: str, *, method: HttpMethod, **kwargs) -> Response:
        with httpx.Client() as client:
            headers = self._default_headers | kwargs.pop("headers", {})
            resp = client.request(
                method=method.value, url=url, headers=headers, **kwargs
            )
            resp.raise_for_status()

        return resp

    @property
    def version(self) -> str:
        server_identity = self.get_server_identity_sync()
        return server_identity.media_container.version
