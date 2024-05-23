from pydantic_settings import BaseSettings, SettingsConfigDict

from gobble.plexapi.config import PlexSettings


class GobbleSettings(BaseSettings):
    plex: PlexSettings

    model_config = SettingsConfigDict(
        env_nested_delimiter="__",
    )


settings = GobbleSettings(_env_file=".env")
