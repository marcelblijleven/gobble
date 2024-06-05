from typing import Generator
from unittest.mock import MagicMock, patch

import pytest

from gobble.homeassistant.client import HomeassistantClient


@pytest.fixture
def mock_homeassistant_client() -> Generator[HomeassistantClient, None, None]:
    mock_client = MagicMock(spec=HomeassistantClient)
    with patch(
        "gobble.homeassistant.client.get_homeassistant_client", return_value=mock_client
    ):
        yield mock_client
