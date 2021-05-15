import logging
from concurrent.futures import ThreadPoolExecutor

from executor.registry import _plugin_registry
from model.vuln import WebVuln, ServiceVuln, Statistics

logger = logging.getLogger("executor")

pool = ThreadPoolExecutor()


def pool_task_wrapper(method, *args):
    try:
        method(*args)
    except Exception as e:
        logger.exception(e)


def dispatch_web_vuln(instance: str, vuln: WebVuln):
    for p in _plugin_registry:
        pool.submit(pool_task_wrapper, p.process_web_vuln, instance, vuln)
        logger.info("submit web vuln call plugin %s" % p.meta.key)


def dispatch_service_vuln(instance: str, vuln: ServiceVuln):
    for p in _plugin_registry:
        pool.submit(pool_task_wrapper, p.process_service_vuln, instance, vuln)
        logger.info("submit service vuln call plugin %s" % p.meta.key)


def dispatch_statistics(instance: str, s: Statistics):
    for p in _plugin_registry:
        pool.submit(pool_task_wrapper, p.process_statistics, instance, s)
        logger.info("submit statistics plugin %s" % p.meta.key)
