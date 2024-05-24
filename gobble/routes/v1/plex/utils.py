from gobble.exceptions import UnsupportedEventTypeException
from gobble.routes.v1.plex.models import WebhookEventModel
from gobble.tasks.webhooks.event_types import EventType


def get_event_type(event: WebhookEventModel) -> EventType:
    """
    Determines the event type for the provided event.

    Args:
        event: the Plex webhook event

    Raises:
         UnsupportedEventTypeException: if the provided event has unsupported event type.

    Returns:
        The event type that matches the provided event.
    """

    match event.event:
        case "media.pause":
            return EventType.MediaPause
        case "media.":
            return EventType.MediaPlay
        case "media.":
            return EventType.MediaRate
        case "media.":
            return EventType.MediaResume
        case "media.":
            return EventType.MediaScrobble
        case "media.":
            return EventType.MediaStop

    raise UnsupportedEventTypeException(event.event, "Plex")