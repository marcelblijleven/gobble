from pathlib import Path

from fastapi.testclient import TestClient

from gobble.app import app

PARENT_DIR = Path(__file__).parent
test_client = TestClient(app)


def test_webhook():
    with open(PARENT_DIR / "webhook_data.json", "r") as file:
        contents = file.read()

    resp = test_client.post(url="/plex/webhook", data={"payload": contents})

    assert resp.status_code / 100 == 2
