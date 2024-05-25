from gobble.exceptions import ServerNotIdentifiedException
from gobble.plex.api.client import PlexClient
from gobble.plex.config import PlexSettings


class PlexServer:
    def __init__(self, settings: PlexSettings) -> None:
        self._settings = settings
        self._api_client = PlexClient(settings)
        self._identifier: str | None = None
        self._name = settings.name
        self._version: str | None = None

    @property
    def identifier(self) -> str:
        """
        The machine identifier of the Plex server

        Returns:
            A string identifier
        """
        if self._identifier is None:
            raise ServerNotIdentifiedException()

        return self._identifier

    @property
    def name(self) -> str:
        """
        The user defined name of the server

        Returns:
            The name of the server
        """
        return self._name

    @property
    def version(self) -> str:
        """
        The version of the Plex server.

        Notes:
            This is not fetched from the server everytime, only when calling the identify method.

        Returns:
            The version string
        """
        if self._version is None:
            raise ServerNotIdentifiedException()

        return self._version

    async def identify(self) -> str:
        """
        Retrieves the machine identifier for the Plex server

        Returns:
            A string identifier
        """
        identity = await self._api_client.get_server_identity()
        self._identifier = identity.media_container.machine_identifier
        self._version = identity.media_container.version

        return self._identifier
