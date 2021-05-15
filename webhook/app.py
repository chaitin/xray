import logging

from flask import Flask, request, redirect

from config import parse_config, get_config
from executor.registry import init_plugin
from views.views import process_web_vuln, process_statistics, process_host_vuln

app = Flask(__name__)

logging.basicConfig(format="[%(levelname)s] %(asctime)s %(name)s %(message)s",
                    datefmt="%Y-%m-%d %H:%M:%S", level=logging.DEBUG)


# 比如可以给一个界面管理当前的插件，看到插件的数据等？
@app.route("/", methods=["GET"])
def index():
    return redirect("https://xray.cool/xray/#/api/api")

@app.route("/webhook", methods=["POST"])
def webhook():
    token = get_config().server_config.token
    if token != "":
        if token != request.args.get("token", ""):
            return "invalid token", 401
    # 可以使用 instance query 来区分不同的节点的数据
    instance = request.args.get("instance", "default")
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
