from pydantic import Field
from pydantic_settings import (
    BaseSettings,
    PydanticBaseSettingsSource,
    SettingsConfigDict,
    TomlConfigSettingsSource,
    YamlConfigSettingsSource,
)

from gobble.discord.config import DiscordSettings
from gobble.plex.config import PlexSettings


class Settings(BaseSettings):
    plex: list[PlexSettings] | None = Field(
        None, description="Settings for one or more Plex servers"
    )
    discord: list[DiscordSettings] | None = Field(
        None, description="Settings for one or more Discord webhooks"
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
