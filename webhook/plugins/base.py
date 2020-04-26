import logging

from cached_property import cached_property

from model.plugin import PluginMeta
from model.vuln import Statistics, WebVuln, ServiceVuln


class BasePlugin:
    @cached_property
    def meta(self) -> PluginMeta:
        raise NotImplementedError("you should implement this method in your subclass")

    @cached_property
    def logger(self):
        return logging.getLogger("plugin:" + self.meta.key)

    def process_web_vuln(self, instance: str, vuln: WebVuln):
        pass

    def process_service_vuln(self, instance: str, vuln: ServiceVuln):
        pass

    def process_statistics(self, instance: str, statistics: Statistics):
        pass
