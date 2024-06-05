from uuid import uuid4

import pytest
from httpx import Response

from gobble.homeassistant.client import get_homeassistant_client


@pytest.mark.asyncio
async def test_client__call_service(respx_mock):
    url = "http://ha.test/"
    token = str(uuid4())
    domain = "light"
    service = "turn_on"
    entity_id = str(uuid4())
    client = get_homeassistant_client(token=token, url=url)

    respx_mock.post(
        f"{url}api/services/{domain}/{service}",
        headers={"Authorization": f"Bearer {token}"},
        json={"entity_id": entity_id},
    ).mock(Response(status_code=200, json={"test": True}))

    resp = await client.call_service(entity_id, domain, service)

    assert resp == {"test": True}
