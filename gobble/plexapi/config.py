from pydantic import AnyHttpUrl, BaseModel, Field


class PlexSettings(BaseModel):
    server_url: AnyHttpUrl
    token: str = Field(..., description="X-Plex-Token")
