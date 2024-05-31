import logging

import httpx
from httpx import HTTPError

from gobble.discord.exceptions import DiscordWebhookException
from gobble.discord.models import DiscordWebhookBody

logger = logging.getLogger(__name__)


async def call_webhook(url: str, body: DiscordWebhookBody) -> None:
    async with httpx.AsyncClient(
        headers={
            "Content-Type": "application/json",
        }
    ) as client:
        try:
            resp = await client.post(
                url, json=body.model_dump(exclude_none=True, exclude_unset=True)
            )
            resp.raise_for_status()
        except HTTPError as exc:
            logger.error(resp.content)
            raise DiscordWebhookException from exc
