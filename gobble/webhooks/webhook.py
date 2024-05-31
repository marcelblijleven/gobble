from enum import Enum

from pydantic import BaseModel, Field

from gobble.webhooks.event_types import EventType


class Source(Enum):
    Plex = "plex"


class MediaType(Enum):
    Show = "show"
    Movie = "movie"
    Music = "music"


class WebhookEvent(BaseModel):
    event: EventType = Field(...)
    source: Source = Field(...)
    media_type: MediaType = Field(...)
    username: str = Field(...)
    title: str = Field(...)
    parent_title: str | None = Field(None)
    grandparent_title: str | None = Field(None)
    index: int | None = Field(...)
    parent_index: int | None = Field(...)
    tvdb: str | None = Field(None)
    imdb: str | None = Field(None)
    tmdb: str | None = Field(None)
