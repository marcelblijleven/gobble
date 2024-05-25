from pydantic import BaseModel, ConfigDict, Field
from pydantic.alias_generators import to_camel


class ServerCapabilitiesMediaContainer(BaseModel):
    size: int
    allow_camera_upload: bool
    allow_channel_access: bool
    allow_media_deletion: bool
    allow_sharing: bool
    allow_sync: bool
    allow_tuners: bool
    background_processing: bool
    certificate: bool
    companion_proxy: bool
    country_code: str
    diagnostics: str
    event_stream: bool
    friendly_name: str
    hub_search: bool
    item_clusters: bool
    livetv: int
    machine_identifier: str
    media_providers: bool
    multiuser: bool
    music_analysis: int
    my_plex: bool
    my_plex_mapping_state: str
    my_plex_signin_state: str
    my_plex_subscription: bool
    my_plex_username: str
    offline_transcode: int
    owner_features: str
    platform: str
    platform_version: str
    plugin_host: bool
    push_notifications: bool
    read_only_libraries: bool
    streaming_brain_abr_version: int = Field(..., alias="streamingBrainABRVersion")
    streaming_brain_version: int
    sync: bool
    transcoder_active_video_sessions: int
    transcoder_audio: bool
    transcoder_lyrics: bool
    transcoder_photo: bool
    transcoder_subtitles: bool
    transcoder_video: bool
    transcoder_video_bitrates: str
    transcoder_video_qualities: str
    transcoder_video_resolutions: str
    updated_at: int
    updater: bool
    version: str
    voice_search: bool
    directory: list[dict] = Field(..., alias="Directory")

    model_config = ConfigDict(alias_generator=to_camel, extra="ignore")


class ServerIdentityMediaContainer(BaseModel):
    size: int = Field(..., alias="size")
    claimed: bool = Field(..., alias="claimed")
    machine_identifier: str = Field(..., alias="machineIdentifier")
    version: str = Field(..., alias="version")


class ServerCapabilitiesResponse(BaseModel):
    media_container: ServerCapabilitiesMediaContainer = Field(
        ..., alias="MediaContainer"
    )


class ServerIdentityResponse(BaseModel):
    media_container: ServerIdentityMediaContainer = Field(..., alias="MediaContainer")
