from fastapi.testclient import TestClient

from gobble.app import app


test_client = TestClient(app)


def test_webhook():
    with open("webhook_data.json", "r") as file:
        contents = file.read()

    resp = test_client.post(url="/plex/webhook", data={"payload": contents})
    ...
