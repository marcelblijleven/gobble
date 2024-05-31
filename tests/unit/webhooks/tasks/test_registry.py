import pytest

from gobble.webhooks.event_types import EventType
from gobble.webhooks.tasks.exceptions import InvalidTaskException
from gobble.webhooks.tasks.registry import autodetect, register_webhook_task
from gobble.webhooks.webhook import WebhookEvent


def test_register_webhook_task__registers_function_to_registry(task_registry):
    assert not task_registry.keys()

    @register_webhook_task(EventType.MediaPlay)
    async def example(_: WebhookEvent): ...

    tasks = task_registry.get(EventType.MediaPlay)
    assert tasks == [example]


def test_register_webhook_task__raises_when_task_is_not_async(task_registry):
    with pytest.raises(Exception) as exc_info:

        @register_webhook_task(EventType.MediaPlay)
        def example(_: WebhookEvent): ...

    assert exc_info.type == InvalidTaskException


def test_autodetect__fills_registry_when_run(task_registry):
    assert not task_registry

    autodetect()

    assert task_registry
