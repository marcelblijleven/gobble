"""
Documentation available at https://discord.com/developers/docs/resources/channel#create-message
"""

from datetime import datetime

from pydantic import BaseModel, Field


class EmbedFooter(BaseModel):
    text: str = Field(...)
    icon_url: str | None = Field(None)
    proxy_icon_url: str | None = None


class EmbedMedia(BaseModel):
    url: str = Field(...)
    proxy_url: str | None = Field(None)
    height: int | None = Field(None)
    width: int | None = Field(None)


class EmbedImage(EmbedMedia): ...


class EmbedThumbnail(EmbedMedia): ...


class EmbedVideo(EmbedMedia): ...


class EmbedProvider(BaseModel):
    name: str | None = Field(None)
    url: str | None = Field(None)


class EmbedAuthor(BaseModel):
    name: str = Field(...)
    url: str | None = Field(None)
    icon_url: str | None = Field(None)
    proxy_icon_url: str | None = Field(None)


class EmbedField(BaseModel):
    name: str = Field(...)
    value: str = Field(...)
    inline: bool | None = Field(None)


class EmbedObject(BaseModel):
    title: str | None = Field(None)
    type: str | None = Field(None)
    description: str | None = Field(None)
    url: str | None = Field(None)
    timestamp: datetime | None = Field(None)
    color: int | None = Field(None)
    footer: EmbedFooter | None = Field(None)
    image: EmbedImage | None = Field(None)
    thumbnail: EmbedThumbnail | None = Field(None)
    video: EmbedVideo | None = Field(None)
    provider: EmbedProvider | None = Field(None)
    author: EmbedAuthor | None = Field(None)
    fields: list[EmbedField] | None = Field(None, max_length=25)


class DiscordWebhookBody(BaseModel):
    content: str
    username: str | None = Field(None)
    avatar_url: str | None = Field(None)
    tts: bool | None = Field(None)
    embeds: list[EmbedObject] | None = Field(None, max_length=10)
