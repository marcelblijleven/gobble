from pydantic import BaseModel, Field, AnyHttpUrl, IPvAnyAddress, ConfigDict


class VersionResponseModel(BaseModel):
    version: str


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


class MetadataModel(BaseModel):
    model_config = ConfigDict(extra="allow")


class WebhookEventModel(BaseModel):
    event: str = Field(...)
    user: bool = Field(...)
    owner: bool = Field(...)
    account: AccountModel = Field(..., alias="Account")
    server: ServerModel = Field(..., alias="Server")
    player: PlayerModel = Field(..., alias="Player")
    metadata: MetadataModel = Field(..., alias="Metadata")
