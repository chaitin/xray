# 对接 xray 和微信实现自动告警

!> 本文的代码写法随着 xray 和第三方服务的升级改进会逐渐失效，请仅参考思路。在 xray Github 仓库 webhook_demo 目录中有最新的实现代码。

## xray 是什么
xray 是从长亭洞鉴核心引擎中提取出的社区版漏洞扫描神器，支持主动、被动多种扫描方式，自备盲打平台、可以灵活定义 POC，功能丰富，调用简单，支持 Windows / macOS / Linux 多种操作系统，可以满足广大安全从业者的自动化 Web 漏洞探测需求。

## 如何第一时间知道扫出了漏洞

对于安全工程师来说，扫描器发现了漏洞能第一时间给出告警是非常重要的，因为安全工程师使用的是 xray 的基础爬虫模式，爬虫一直在爬也不会一直人工刷新和查看漏洞报告，也有可能是使用的被动代理模式，让测试人员挂扫描器代理然后访问各个业务页面，但是不知道什么时间测试人员才开始和完成测试，也有可能是日志扫描模式，导入日志使用脚本进行 url 扫描，不知道什么时间才能重放完成。

还有很多公司自建了漏洞管理系统、工单系统等等，如果扫描器发现了漏洞可以自动同步这些系统也将会大大解放安全人员。针对这些场景 xray 有一种漏洞输出模式叫 `webhook-output`，在发现漏洞的时候，将会向指定的 url post 漏洞数据，demo 的代码就是 

```python
import requests
requests.post(webhook, json=vuln_info)
```

如果我们写一个中间的转换和转发层，就可以很方便的实现下面的功能了

 - 发送邮件、短信告警
 - 发送微信、企业微信、钉钉、slack告警
 - 漏洞信息同步到自己的数据库中
 - 为该漏洞创建一个工单
 - 使用其他的工具去验证漏洞是否存在
 - ...... 

## 使用 webhook 做自动推送
本文就借助 [Server酱](http://sc.ftqq.com/3.version) 和[企业微信机器人](https://work.weixin.qq.com/help?person_id=1&doc_id=13376)，来演示如何实时通知 xray 发现了漏洞。

## xray 的 webhook 是什么
对于 xray，webhook 应该是一个 url 地址，也就是我们需要自己搭建一个 web 服务器，接收到 xray 发送的漏洞信息，然后在将它转发，借助于 Python 的 flask 框架，我们很快写了一个 webhook url 的 demo 出来。

```
from flask import Flask, request
import requests

app = Flask(__name__)

@app.route('/webhook', methods=['POST'])
def xray_webhook():
    print(request.json)
    return 'ok'

if __name__ == '__main__':
    app.run()
```

使用 `xray webscan --url http://pentester-web.vulnet/sqli/example1.php?name=root --plugins sqldet --webhook-output http://127.0.0.1:5000/webhook` 测试，然后发现成功打印出了漏洞信息。
 
```shell
 * Serving Flask app "app.py"
 * Environment: development
 * Debug mode: off
 * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
{'create_time': 1566836256580, 'detail': {'host': 'pentester-web.vulnet', 'param': {'key': 'name', 'position': 'query', 'value': "root'and'lW'='lql"}, 'payload': "root'and'lW'='lql", 'port': 80, 'request': '', 'request1': 'GET /sqli/example1.php?name=root%27and%274w%27%3D%274w HTTP/1.1\r\nHost: pentester-web.vulnet\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169\r\nCookie: key=value\r\nAccept-Encoding: gzip\r\n\r\n', 'request2': 'GET /sqli/example1.php?name=root%27and%27lW%27%3D%27lql HTTP/1.1\r\nHost: pentester-web.vulnet\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169\r\nCookie: key=value\r\nAccept-Encoding: gzip\r\n\r\n', 'response': '', 'response1': 'HTTP/1.1 200 OK\r\n...', 'response2': 'HTTP/1.1 200 OK\r\n...', 'title': "Generic Boolean based case ['string']", 'type': 'boolean_based', 'url': 'http://pentester-web.vulnet/sqli/example1.php?name=root'}, 'plugin': 'sqldet', 'target': {'url': 'http://pentester-web.vulnet/sqli/example1.php', 'params': [{'position': 'query', 'path': ['name']}]}, 'vuln_class': ''}
127.0.0.1 - - [27/Aug/2019 00:17:36] "POST /webhook HTTP/1.1" 200 -
```

接下来就是解析 xray 的漏洞信息，然后生成对应的页面模板就好了。需要参考[文档](/api/vuln)。因为推送不适合发送太大的数据量，所以就选择了基础的一些字段。

```python
from flask import Flask, request
import requests

app = Flask(__name__)


@app.route('/webhook', methods=['POST'])
def xray_webhook():
    vuln = request.json
    # 因为还会收到 https://chaitin.github.io/xray/#/api/statistic 的数据
    if "vuln_class" not in vuln:
        return "ok"
    content = """## xray 发现了新漏洞
    
url: {url}

插件: {plugin}

漏洞类型: {vuln_class}

发现时间: {create_time}

请及时查看和处理
""".format(url=vuln["target"]["url"], plugin=vuln["plugin"],
           vuln_class=vuln["vuln_class"] or "Default",
           create_time=str(datetime.datetime.fromtimestamp(vuln["create_time"] / 1000)))
    print(content)
    return 'ok'

if __name__ == '__main__':
    app.run()
```
 
### Server 酱

Server酱是一款程序员和服务器之间的通信软件，也就是从服务器推报警和日志到手机的工具。

开通并使用上它还是很简单的

 - 登入：用 GitHub 账号登录 [http://sc.ftqq.com/3.version](http://sc.ftqq.com/3.version)，就能获得一个 SECKEY 
 - 绑定：扫码关注完成绑定
 - 发消息：往 `http://sc.ftqq.com/{SECKEY}.send` 发请求，就可以在微信里收到消息啦

 我们先用 Python 写一个简单的 demo，以下所有的 SECKEY 的实际值我都使用 `{SECKEY}` 代替，大家需要修改为自己的值。
 
```python
import requests
requests.post("https://sc.ftqq.com/{SECKEY}.send", 
              data={"text": "xray vuln alarm", "desp": "test content"})
```

很简单就收到了消息，将上面 xray 的漏洞信息结合在一起，就是

```python
from flask import Flask, request
import requests
import datetime
import logging

app = Flask(__name__)


def push_ftqq(content):
    resp = requests.post("https://sc.ftqq.com/{SECKEY}.send",
                  data={"text": "xray vuln alarm", "desp": content})
    if resp.json()["errno"] != 0:
        raise ValueError("push ftqq failed, %s" % resp.text)

@app.route('/webhook', methods=['POST'])
def xray_webhook():
    vuln = request.json
    # 因为还会收到 https://chaitin.github.io/xray/#/api/statistic 的数据
    if "vuln_class" not in vuln:
        return "ok"
    content = """## xray 发现了新漏洞
    
url: {url}

插件: {plugin}

漏洞类型: {vuln_class}

发现时间: {create_time}

请及时查看和处理
""".format(url=vuln["target"]["url"], plugin=vuln["plugin"],
           vuln_class=vuln["vuln_class"] or "Default",
           create_time=str(datetime.datetime.fromtimestamp(vuln["create_time"] / 1000)))
    try:
        push_ftqq(content)
    except Exception as e:
        logging.exception(e)
    return 'ok'


if __name__ == '__main__':
    app.run()
```

展示效果如图

![](../assets/scenario/xray_vuln_alert/1.jpg)


### 企业微信群机器人

企业微信群机器人就像一个普通成员一样，可以发言，可以 `@` 人，如果我们接入企业微信群做 xray 的漏洞告警，也会大大方便漏洞的第一时间发现。

开通和使用方法

 - 点击群聊右上角，然后找到 '群机器人'，然后点击'添加'
 - 复制 Webhook 的地址，保存备用

 ![](../assets/scenario/xray_vuln_alert/2.jpg)


调用的代码也非常简单，我们只需要展示主要的部分就可以了

```python
def push_wechat_group(content):
    resp = requests.post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key={KEY}",
                         json={"msgtype": "markdown",
                               "markdown": {"content": content}})
    if resp.json()["errno"] != 0:
        raise ValueError("push wechat group failed, %s" % resp.text)
```

展示效果如图

![](../assets/scenario/xray_vuln_alert/3.jpg)
