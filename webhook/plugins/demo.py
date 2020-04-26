import time

from cached_property import cached_property

from model.plugin import PluginMeta
from model.vuln import Statistics, WebVuln, ServiceVuln
from plugins.base import BasePlugin


class Plugin(BasePlugin):
    def __init__(self, arg1, arg2):
        self.arg1 = arg1
        self.arg2 = arg2

    @cached_property
    def meta(self):
        return PluginMeta(key="DemoPlugin", description="hello world", author="chaitin")

    def process_statistics(self, instance: str, statistics: Statistics):
        self.logger.info(statistics)

    def process_service_vuln(self, instance: str, vuln: ServiceVuln):
        self.logger.info(vuln)

    def process_web_vuln(self, instance: str, vuln: WebVuln):
        time.sleep(2)
        self.logger.info("new vuln instance: %s, plugin: %s, url: %s" % (instance, vuln.plugin, vuln.url))
