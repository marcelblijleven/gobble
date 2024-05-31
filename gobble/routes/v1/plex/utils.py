from gobble.exceptions import UnsupportedEventTypeException
from gobble.routes.v1.plex.models import PlexWebhookEventModel
from gobble.webhooks.event_types import EventType


def get_event_type(event: PlexWebhookEventModel) -> EventType:
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
        case "media.play":
            return EventType.MediaPlay
        case "media.rate":
            return EventType.MediaRate
        case "media.resume":
            return EventType.MediaResume
        case "media.scrobble":
            return EventType.MediaScrobble
        case "media.stop":
            return EventType.MediaStop

    server_ident = f"Plex ({event.server.title})"
    raise UnsupportedEventTypeException(event.event, server_ident)
