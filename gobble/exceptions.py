class UnsupportedEventTypeException(Exception):
    def __init__(self, event_type: str, source: str) -> None:
        msg = f"received an unsupported event type '{event_type}' from {source}"
        super().__init__(msg)


class ServerIdentificationException(Exception): ...


class ServerNotIdentifiedException(ServerIdentificationException): ...
