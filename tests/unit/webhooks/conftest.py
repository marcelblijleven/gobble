from typing import Generator
from unittest.mock import patch

import pytest

from gobble.webhooks.tasks.registry import WebhookTaskRegistry


@pytest.fixture
def task_registry() -> Generator[WebhookTaskRegistry, None, None]:
    with patch(
        "gobble.webhooks.tasks.registry.webhook_task_registry",
        new=WebhookTaskRegistry(),
    ) as patched_registry:
        yield patched_registry
