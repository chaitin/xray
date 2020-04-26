import datetime
import logging
from datetime import datetime

from flask import Flask, request, redirect

from config import parse_config
from executor.executor import dispatch_web_vuln, dispatch_service_vuln, dispatch_statistics
from executor.registry import init_plugin
from model.vuln import Statistics, WebVuln, WebParam, WebParamPosition, WebRequest, WebResponse, ServiceVuln

app = Flask(__name__)

debug = True
logging.basicConfig(format="[%(levelname)s] %(asctime)s %(name)s %(message)s",
                    datefmt="%Y-%m-%d %H:%M:%S", level=logging.DEBUG)

logger = logging.getLogger("webhook")


def process_web_vuln(instance, data):
    # 数据格式见 https://xray.cool/xray/#/api/vuln
    detail = data["detail"]
    p = detail["param"]
    if p:
        param = WebParam(key=p["key"], value=p["value"], position=WebParamPosition(p["position"]))
    else:
        param = None

    request = []
    response = []
    extra = {}

    for i in range(0, 10):
        req_key = f"request{i}" if i else "request"
        resp_key = f"response{i}" if i else "response"
        req = detail.get(req_key)
        resp = detail.get(resp_key)

        if req == "" or resp == "":
            continue
        if req is None or resp is None:
            break
        request.append(WebRequest(raw=req))
        response.append(WebResponse(raw=resp))

    not_extra_key = ["request", "response", "param", "payload", "url"]
    for k, v in detail.items():
        for item in not_extra_key:
            if item in k:
                break
        else:
            extra[k] = v

    vuln = WebVuln(create_time=datetime.fromtimestamp(data["create_time"] / 1000), plugin=data["plugin"],
                   vuln_class=data["vuln_class"],
                   url=data["target"]["url"], param=param, request=request, response=response, extra=extra)
    dispatch_web_vuln(instance, vuln)


def process_statistics(instance, data):
    data.pop("type", None)
    # 数据格式见 https://xray.cool/xray/#/api/statistic
    s = Statistics(**data)
    dispatch_statistics(instance, s)


def process_host_vuln(instance, data):
    detail = data["detail"]
    extra = {}

    not_extra_key = ["host", "port"]
    for k, v in detail.items():
        for item in not_extra_key:
            if item in k:
                break
        else:
            extra[k] = v

    vuln = ServiceVuln(create_time=datetime.fromtimestamp(data["create_time"] / 1000), plugin=data["plugin"],
                       vuln_class=data["vuln_class"], host=detail["host"], port=detail["port"],
                       extra=extra)
    dispatch_service_vuln(instance, vuln)


# 比如可以给一个界面管理当前的插件，看到插件的数据等？
@app.route("/", methods=["GET"])
def index():
    return redirect("https://xray.cool/xray/#/api/api")


@app.route("/webhook", methods=["POST"])
def webhook():
    # 可以使用 instance query 来区分不同的节点的数据
    instance = request.args.get("instance", "")
    # FIXME 使用 token 来保证 api 安全
    data = request.json
    data_type = data.get("type")
    if data_type == "web_vuln":
        process_web_vuln(instance, data)
    elif data_type == "web_statistic":
        process_statistics(instance, data)
    elif data_type == "host_vuln":
        process_host_vuln(instance, data)
    return "ok"


if __name__ == "__main__":
    config = parse_config("./config.yml")
    init_plugin(config)
    app.run(debug=config.server_config.debug, host=config.server_config.host, port=config.server_config.port)
