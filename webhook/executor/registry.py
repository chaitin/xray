from typing import List

from model.config import Config
from plugins.base import BasePlugin

_plugin_registry: List[BasePlugin] = []


def import_class(cl):
    d = cl.rfind(".")
    class_name = cl[d + 1:len(cl)]
    m = __import__(cl[0:d], globals(), locals(), [class_name])
    return getattr(m, class_name)


def init_plugin(config: Config):
    for cls, conf in config.plugin_config.items():
        if not conf.enabled:
            continue
        if "." not in cls:
            cls = cls + ".Plugin"
        p = import_class("plugins." + cls)
        _plugin_registry.append(p(**conf.args))
