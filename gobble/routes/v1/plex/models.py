import logging

from pydantic import (
    AnyHttpUrl,
    BaseModel,
    ConfigDict,
    Field,
    IPvAnyAddress,
    RootModel,
)

from gobble.webhooks.event_types import EventType
from gobble.webhooks.exceptions import UnknownEventTypeException
from gobble.webhooks.webhook import Source, WebhookEvent

logger = logging.getLogger(__name__)


class VersionResponseModel(RootModel):
    root: dict[str, str]


class AccountModel(BaseModel):
    id: int = Field(...)
    thumb: AnyHttpUrl = Field(...)
    title: str = Field(...)


class ServerModel(BaseModel):
    uuid: str = Field(...)
    title: str = Field(...)


class PlayerModel(BaseModel):
    local: bool = Field(...)
    public_address: IPvAnyAddress = Field(..., alias="publicAddress")
    title: str = Field(...)
    uuid: str = Field(...)


class Guid(BaseModel):
    id: str


class Rating(BaseModel):
    image: str = Field(...)
    value: float = Field(...)
    type: str = Field(...)


class Identity(BaseModel):
    id: int = Field(...)
    filter: str = Field(...)
    tag: str = Field(...)
    tag_key: str = Field(..., alias="tagKey")


class Director(Identity): ...


class Writer(Identity): ...


class Role(Identity):
    role: str = Field(...)
    thumb: str | None = Field(None)


class MetadataModel(BaseModel):
    library_section_type: str = Field(..., alias="librarySectionType")
    rating_key: str = Field(..., alias="ratingKey")
    key: str = Field(..., alias="key")
    parent_rating_key: str | None = Field(None, alias="parentRatingKey")
    grand_parent_rating_key: str | None = Field(None, alias="grandparentRatingKey")
    guid: str = Field(..., alias="guid")
    parent_guid: str | None = Field(None, alias="parentGuid")
    grand_parent_guid: str | None = Field(None, alias="grandparentGuid")
    grand_parent_slug: str | None = Field(None, alias="grandparentSlug")
    type: str = Field(..., alias="type")
    title: str = Field(..., alias="title")
    grand_parent_key: str | None = Field(None, alias="grandparentKey")
    parent_key: str | None = Field(None, alias="parentKey")
    library_section_title: str = Field(..., alias="librarySectionTitle")
    library_section_id: int = Field(..., alias="librarySectionID")
    library_section_key: str = Field(..., alias="librarySectionKey")
    grand_parent_title: str | None = Field(None, alias="grandparentTitle")
    parent_title: str | None = Field(None, alias="parentTitle")
    content_rating: str | None = Field(None, alias="contentRating")
    summary: str = Field(..., alias="summary")
    index: int | None = Field(None, alias="index")
    parent_index: int | None = Field(None, alias="parentIndex")
    audience_rating: float | None = Field(None, alias="audienceRating")
    view_offset: int | None = Field(None, alias="viewOffset")
    last_viewed_at: int | None = Field(None, alias="last_viewed_at")
    year: int | None = Field(None, alias="year")
    thumb: str = Field(..., alias="thumb")
    art: str = Field(..., alias="art")
    parent_thumb: str | None = Field(None, alias="parentThumb")
    grandparent_thumb: str | None = Field(None, alias="grandparentThumb")
    grandparent_art: str | None = Field(None, alias="grandparentArt")
    grandparent_theme: str | None = Field(None, alias="grandparentTheme")
    duration: int | None = Field(None, alias="duration")
    originally_available_at: str | None = Field(None, alias="originallyAvailableAt")
    added_at: int = Field(..., alias="addedAt")
    updated_at: int = Field(..., alias="updatedAt")
    audience_rating_image: str | None = Field(None, alias="audienceRatingImage")
    chapter_source: str | None = Field(None, alias="chapterSource")
    guid_list: list[Guid] = Field(..., alias="Guid")
    rating: list[Rating] = Field(default_factory=list, alias="Rating")
    director: list[Director] = Field(default_factory=list, alias="Director")
    writer: list[Writer] = Field(default_factory=list, alias="Writer")
    role: list[Role] = Field(default_factory=list, alias="Role")


class PlexWebhookEventModel(BaseModel):
    event: str = Field(...)
    user: bool = Field(...)
    owner: bool = Field(...)
    account: AccountModel = Field(..., alias="Account")
    server: ServerModel = Field(..., alias="Server")
    player: PlayerModel = Field(..., alias="Player")
    metadata: MetadataModel = Field(..., alias="Metadata")

    model_config = ConfigDict(extra="allow")

    def to_generic_webhook_event(self) -> WebhookEvent:
        """
        Converts the Plex Webhook event to a generic webhook event

        Returns:
            WebhookEvent
        """
        tvdb: str | None = None
        imdb: str | None = None
        tmdb: str | None = None

        for guid in self.metadata.guid_list:
            if guid.id.startswith("tvdb"):
                tvdb = guid.id
            if guid.id.startswith("imdb"):
                imdb = guid.id
            if guid.id.startswith("tmdb"):
                tmdb = guid.id

        media_type = self.metadata.library_section_type
        media_type = "music" if media_type == "artist" else media_type

        return WebhookEvent(
            event=_plex_event_to_event_type(self.event),
            source=Source.Plex,
            media_type=media_type,
            username=self.account.title,
            title=self.metadata.title,
            parent_title=self.metadata.parent_title,
            grandparent_title=self.metadata.grand_parent_title,
            index=self.metadata.index,
            parent_index=self.metadata.parent_index,
            tvdb=tvdb,
            imdb=imdb,
            tmdb=tmdb,
        )


def _plex_event_to_event_type(event: str) -> EventType:
    """
    Converts a Plex event to an EventType

    Args:
        event: the Plex webhook event

    Returns:
        EventType value
    """
    try:
        event = event.replace(".", "_")
        event_type = EventType(event)
    except ValueError as exc:
        logger.error(f"received unknown event type: {event}")
        raise UnknownEventTypeException from exc

    return event_type
