from dataclasses import dataclass
from typing import Any, Dict


@dataclass
class ServerConfig:
    host: str
    port: int
    debug: bool
    token: str
    base_url: str


@dataclass
class PluginConfig:
    enabled: bool
    args: Dict[str, Any]


@dataclass
class Config:
    version: int
    server_config: ServerConfig
    plugin_config: Dict[str, PluginConfig]
