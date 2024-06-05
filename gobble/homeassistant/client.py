from functools import cache
from typing import Any

import httpx


class HomeassistantClient:
    def __init__(self, token: str, url: str) -> None:
        self._token = token
        self._url = url

    async def call_service(
        self, entity_id: str, domain: str, service: str, data: dict | None = None
    ):
        """
        Trigger a specific domain service for the provided entity id.

        Args:
            entity_id: the entity id that the service should run on
            domain: the domain of the service to target
            service: the actual service
            data: optional additional data to call the service with

        Returns:
            The response body
        """

        body: dict[str, Any] = {"entity_id": entity_id}

        if data:
            body.update({"data": data})

        async with httpx.AsyncClient() as client:
            resp = await client.post(
                f"{self._url}api/services/{domain}/{service}",
                headers={"Authorization": f"Bearer {self._token}"},
                json=body,
            )

            resp.raise_for_status()
            return resp.json()


@cache
def get_homeassistant_client(token: str, url: str) -> HomeassistantClient:
    return HomeassistantClient(token, url)
