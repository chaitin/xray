from dataclasses import dataclass
from datetime import datetime
from enum import Enum
from typing import List, Any, Dict, Optional


@dataclass
class Statistics:
    """统计数据"""
    raw_json: dict
    # 发现的 url 数量
    num_found_urls: int
    # 扫描完成的 url 数量
    # num_found_urls - num_scanned_urls 就是还没有扫描的请求数量
    num_scanned_urls: int
    # 扫描发送的 http 请求数量
    num_sent_http_requests: int
    # 最近 30s 内平均响应时间
    average_response_time: float
    # 最近 30s 请求失败率
    ratio_failed_http_requests: float
    # 暂时无用字段
    ratio_progress: float


class WebParamPosition(Enum):
    """参数位置"""
    query = "query"
    body = "body"
    cookie = "cookie"
    header = "header"


@dataclass
class WebParam:
    # 参数名
    key: str
    # 参数值
    value: str
    # 参数位置
    position: WebParamPosition


@dataclass
class WebRequest:
    # http 原始请求
    raw: str


@dataclass
class WebResponse:
    # http 原始响应
    raw: str


@dataclass
class WebVuln:
    """web 漏洞"""
    raw_json: dict
    # 创建时间
    create_time: datetime
    # 这两个数据内部使用其实是 enum，要不要提供给社区？
    # 插件名
    plugin: str
    # 漏洞类型，可能为空，代表 default
    vuln_class: str

    url: str
    # 存在漏洞的参数，可能为 None
    param: Optional[WebParam]
    # 证明漏洞存在的请求序列
    request: List[WebRequest]
    # 证明漏洞存在的响应序列
    response: List[WebResponse]
    # 插件开发者可以添加各种额外数据，比如作者名、证明漏洞存在的其他数据等
    # 在 xray 内部，可能是 map[string]string、map[string][]string 和 map[string]map[string]string
    extra: Dict[str, Any]


@dataclass
class ServiceVuln:
    """服务漏洞"""
    raw_json: dict
    # 同 web 漏洞
    create_time: datetime
    plugin: str
    vuln_class: str

    # 主机名
    host: str
    # 端口
    port: int
    # 同 web 漏洞
    extra: Dict[str, Any]
