import asyncio
import importlib.util
import inspect
import logging.config
from pathlib import Path
from typing import (
    Any,
    Callable,
    Coroutine,
    ParamSpec,
    TypeAlias,
    TypeVar,
)

from gobble.logging_config import LOGGING_CONFIG
from gobble.webhooks.event_types import EventType
from gobble.webhooks.tasks.exceptions import InvalidTaskException
from gobble.webhooks.webhook import WebhookEvent

_TASKS_DIR = Path(__file__).parent
P = ParamSpec("P")
R = TypeVar("R")
C: TypeAlias = Coroutine[Any, Any, R]
RegistryDictType: TypeAlias = dict[EventType, list[Callable[P, C]]]

logging.config.dictConfig(LOGGING_CONFIG)
logger = logging.getLogger(__name__)


class WebhookTaskRegistry:
    def __init__(self, register_to: RegistryDictType | None = None):
        if register_to is None:
            register_to = {}
        self._registry: RegistryDictType = register_to
        self._task_names: set[str] = set()

    @property
    def task_names(self) -> list[str]:
        """
        Returns the names of the registered task functions

        Returns:
            list of str
        """
        return list(self._task_names)

    def get(
        self, key: EventType, default: list | None = None
    ) -> list[Callable[[Any], Coroutine[Any, Any, R]]] | None:
        """
        Retrieves a value from the registry by key. If the
        key does not exist, None will be returned.

        Args:
            key: the event type to use as key
            default: default value if key is not found

        Returns:
            A registered task if the key exists, else None
        """

        try:
            return self.__getitem__(key)
        except KeyError:
            return default

    def keys(self):
        """
        Retrieve the registered keys from the registry

        Returns:
            dict_keys
        """
        return self._registry.keys()

    def values(self):
        """
        Retrieve the registered values from the registry

        Returns:
            dict_values
        """
        return self._registry.values()

    def __bool__(self):
        return bool(self._registry)

    def __contains__(self, item):
        return item in self._registry

    def __getitem__(self, item):
        value = self._registry[item]
        return value

    def __setitem__(self, key, value):
        self._registry[key] = value
        return None


webhook_task_registry = WebhookTaskRegistry()


def register_webhook_task(
    *event_type: EventType,
) -> Callable[[Callable[P, C]], Callable[P, C]]:
    """
    Registers a task to be called

    Args:
        *event_type: one or more event types to subscribe to

    Raises:
        InvalidTaskException: when the decorator function is not async

    Returns:
        A task decorator
    """

    def register(func: Callable[P, C]) -> Callable[P, C]:
        if not inspect.iscoroutinefunction(func):
            raise InvalidTaskException("the provided task function is not async")
        for type_ in event_type:
            logger.info(
                f"Registering task '{func.__name__}' for event type '{type_.value}'"
            )
            if type_ in webhook_task_registry:
                webhook_task_registry[type_].append(func)
            else:
                webhook_task_registry[type_] = [func]

        return func

    return register


def autodetect():
    """
    Automatically imports all python files in the webhooks.tasks module
    to register them into the registry

    Returns:
        None
    """
    from gobble.config import settings

    for file in _TASKS_DIR.rglob("*.py"):
        if file.name == Path(__file__).name:
            continue
        if bool(getattr(settings.tasks, file.stem, None)):
            importlib.import_module(f"gobble.webhooks.tasks.{file.stem}")


async def call_registered_webhook_tasks_for_event(event: WebhookEvent) -> None:
    """
    Will create a task group for the registered callbacks that belong to the provided event type.

    Args:
        event: a webhook event model

    Returns:
        None
    """
    event_type = event.event
    async with asyncio.TaskGroup() as task_group:
        if (tasks := webhook_task_registry.get(event_type)) is None:
            logger.info(f"no tasks registered for {event_type=}")
            return

        for task in tasks:
            task_group.create_task(task(event))

    logger.debug(f"finished tasks for {event_type=}")
