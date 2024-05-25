import logging

from gobble.config import Settings
from gobble.exceptions import ServerIdentificationException
from gobble.plex.server import PlexServer
from gobble.protocols import MediaServer

logger = logging.getLogger(__name__)


async def _identify_server(server: MediaServer) -> str:
    """
    Calls the identify method on the provided server.

    Args:
        server: a class instance that conforms to the MediaServer protocol

    Raises:
        ServerIdentificationException: when the server could not be identified or reached

    Returns:
        A string identifier
    """

    try:
        identifier = await server.identify()
    except Exception as exc:
        raise ServerIdentificationException from exc

    return identifier


async def identify_servers(settings: Settings) -> dict[str, MediaServer]:
    """
    For each of the media servers defined in the provided settings, a MediaServer
    instance will be created and its 'identify' method is called.

    If the call is successful, the server is added to a dict with its identifier as
    a key.

    Args:
        settings: the Gobble settings

    Returns:
        A dict of str keys and MediaServer values
    """
    servers: dict[str, MediaServer] = {}

    # Plex servers
    if (plex_settings_list := settings.plex) is not None:
        for plex_settings in plex_settings_list:
            logger.info(f"Attempting to identify Plex server {plex_settings.name}")
            server = PlexServer(plex_settings)

            try:
                identifier = await _identify_server(server)
                servers[identifier] = server
                logger.info(
                    f"Identified Plex server '{plex_settings.name}' with identifier '{server.identifier}'"
                )
            except ServerIdentificationException as exc:
                logger.error(
                    f"Plex server {plex_settings.name} could not be identifier: {exc}"
                )

    return servers
