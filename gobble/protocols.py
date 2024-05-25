from typing import Protocol


class MediaServer(Protocol):
    async def identify(self) -> str:
        """
        Retrieve the identifier for the media server

        Returns:
            A string identifier
        """
        ...

    @property
    def identifier(self) -> str: ...

    @property
    def name(self) -> str: ...

    @property
    def version(self) -> str: ...
