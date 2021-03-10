from pydantic import BaseModel
import typing
import os


class SettingsModel(BaseModel):
    trustix_rpc: str = os.environ["TRUSTIX_RPC"]
    default_attr: str = "hello.x86_64-linux"
    binary_cache_proxy: str = os.environ["TRUSTIX_BINARY_CACHE_PROXY"]
    db_uri: str = os.environ["DB_URI"]

    @property
    def tortoise_config(self) -> typing.Dict:
        return {
            "connections": {
                "default": self.db_uri,
            },
            "apps": {
                "trustix_dash": {
                    "models": ["trustix_dash.models", "aerich.models"],
                }
            },
            "use_tz": False,
            "timezone": "UTC",
        }


settings = SettingsModel()


TORTOISE_ORM = settings.tortoise_config