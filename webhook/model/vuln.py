from dataclasses import dataclass
from datetime import datetime
from enum import Enum
from typing import List, Any, Dict, Optional


@dataclass
class Statistics:
    num_found_urls: int
    num_scanned_urls: int
    num_sent_http_requests: int
    average_response_time: float
    ratio_failed_http_requests: float
    ratio_progress: float


class WebParamPosition(Enum):
    query = "query"
    body = "body"
    cookie = "cookie"
    header = "header"


@dataclass
class WebParam:
    key: str
    value: str
    position: WebParamPosition


@dataclass
class WebRequest:
    raw: str


@dataclass
class WebResponse:
    raw: str


@dataclass
class WebVuln:
    create_time: datetime
    # 这两个数据内部使用其实是 enum，要不要提供给社区？
    plugin: str
    vuln_class: str

    url: str
    param: Optional[WebParam]
    request: List[WebRequest]
    response: List[WebResponse]

    extra: Dict[str, Any]


@dataclass
class ServiceVuln:
    create_time: datetime
    plugin: str
    vuln_class: str

    host: str
    port: int

    extra: Dict[str, Any]
