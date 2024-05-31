from pydantic import BaseModel, Field
from pydantic_settings import (
    BaseSettings,
    PydanticBaseSettingsSource,
    SettingsConfigDict,
    TomlConfigSettingsSource,
    YamlConfigSettingsSource,
)

from gobble.discord.config import DiscordSettings
from gobble.plex.config import PlexSettings


class Tasks(BaseModel):
    logger: bool | None = Field(True)
    discord: list[DiscordSettings] = Field(
        default_factory=list, description="Discord task settings"
    )


class Settings(BaseSettings):
    plex: list[PlexSettings] | None = Field(
        None, description="Settings for one or more Plex servers"
    )
    tasks: Tasks = Field(
        default_factory=Tasks, description="Collection of webhook tasks to register"
    )

    model_config = SettingsConfigDict(yaml_file="gobble.yaml")

    @classmethod
    def settings_customise_sources(
        cls,
        settings_cls: type[BaseSettings],
        init_settings: PydanticBaseSettingsSource,
        env_settings: PydanticBaseSettingsSource,
        dotenv_settings: PydanticBaseSettingsSource,
        file_secret_settings: PydanticBaseSettingsSource,
    ) -> tuple[PydanticBaseSettingsSource, ...]:
        return (
            env_settings,
            init_settings,
            YamlConfigSettingsSource(settings_cls),
            TomlConfigSettingsSource(settings_cls),
        )


def _get_settings() -> Settings:
    return Settings()


settings = _get_settings()
