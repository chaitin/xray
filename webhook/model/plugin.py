from dataclasses import dataclass


@dataclass
class PluginMeta:
    key: str
    description: str
    author: str
