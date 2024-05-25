from pydantic import AnyHttpUrl, BaseModel, Field


class PlexSettings(BaseModel):
    name: str = Field(..., description="Name of the Plex server")
    server_url: AnyHttpUrl = Field(
        ..., description="Url + port your Plex server runs on"
    )
    token: str = Field(..., description="X-Plex-Token")
