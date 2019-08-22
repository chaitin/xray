# 漏洞 json 格式说明

在 `--json-output` 或者 `--webhook-output` 的时候，将会使用 json 格式输出漏洞信息，json 字段说明如下。

## 样例

sql 注入输出样例

```json
{
  "create_time": 1566456018640,
  "detail": {
    "host": "pentester-web.vulnet",
    "param": {
      "key": "name",
      "position": "query",
      "value": "root'and'u1'='zTk"
    },
    "payload": "root'and'u1'='zTk",
    "port": 80,
    "request": "",
    "request1": "GET /sqli/example1.php?name=root%27and%27eF%27%3D%27eF HTTP/1.1\r\nHost: pentester-web.vulnet\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169\r\nCookie: key=value\r\nAccept-Encoding: gzip\r\n\r\n",
    "request2": "GET /sqli/example1.php?name=root%27and%27u1%27%3D%27zTk HTTP/1.1\r\nHost: pentester-web.vulnet\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169\r\nCookie: key=value\r\nAccept-Encoding: gzip\r\n\r\n",
    "response": "",
    "response1": "HTTP/1.1 200 OK\r\n...",
    "response2": "HTTP/1.1 200 OK\r\n...",
    "title": "Generic Boolean based case ['string']",
    "type": "boolean_based",
    "url": "http://pentester-web.vulnet/sqli/example1.php?name=root"
  },
  "plugin": "sqldet",
  "target": {
    "url": "http://pentester-web.vulnet/sqli/example1.php",
    "params": [
      {
        "position": "query",
        "path": [
          "name"
        ]
      }
    ]
  },
  "vuln_class": ""
}

```

baseline 输出样例

```json
{
  "create_time": 1566456530730,
  "detail": {
    "expected_value": "^1; mode=block$",
    "header_name": "X-XSS-Protection",
    "header_value": "0",
    "host": "pentester-web.vulnet",
    "param": {},
    "payload": "",
    "port": 80,
    "request": "GET /sqli/example1.php?name=root HTTP/1.1\r\nHost: pentester-web.vulnet\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169\r\nCookie: key=value\r\nAccept-Encoding: gzip\r\n\r\n",
    "response": "HTTP/1.1 200 OK\r\n...",
    "url": "http://pentester-web.vulnet/sqli/example1.php?name=root"
  },
  "plugin": "baseline",
  "target": {
    "url": "http://pentester-web.vulnet/sqli/example1.php"
  },
  "vuln_class": "header-wrong-value"
}
```

## 字段说明

 - create_time 发现漏洞的时间
 - plugin 插件名
 - target.url 扫描目标 url
 - target.params 历史兼容问题，建议使用 detail.param
 - vuln_class 漏洞类型，部分插件只有一种漏洞类型，可能就是空字符串，代表 default
 - detail 漏洞详情
   - host、url、port 扫描目标信息
   - param 参数信息，部分漏洞无参数
   - payload 部分漏洞无 payload
   - request、response 或者 `request{N}`、`response{N}` 漏洞探测部分输入输出，大部分情况下只有一对 request/response，对于有多组的情况下，比如 sql 布尔注入，会显示两种 payload 的请求响应，就会使用 request1/response1 这种，后面的数字从 1 开始递增，一定是成对的。
   - detail 中其他字段是每个插件自定义的


