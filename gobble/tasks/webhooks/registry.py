from typing import Any, Callable, Coroutine, ParamSpec, TypeAlias, TypeVar

from gobble.tasks.webhooks.event_types import EventType

P = ParamSpec("P")
R = TypeVar("R")
C: TypeAlias = Coroutine[Any, Any, R]

webhook_task_registry: dict[EventType, list[Callable[P, C]]] = {}


def register_webhook_task(
    *event_type: EventType,
) -> Callable[[Callable[P, C]], Callable[P, C]]:
    """
    Registers a task to be called

    :param event_type: the webhook event type
    :return:
    """

    def register(func: Callable[P, C]) -> Callable[P, C]:
        for type_ in event_type:
            if type_ in webhook_task_registry:
                webhook_task_registry[type_].append(func)
            else:
                webhook_task_registry[type_] = [func]

        return func

    return register
