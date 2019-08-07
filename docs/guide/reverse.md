# 反连平台

## 什么是 ssrf，反连平台是干什么的

![](https://chaitin.github.io/xray/assets/reverse.jpg)

## 反连平台配置项目

```yaml
reverse:
  store_events: false
  token: ""
  http:
    enabled: true
    listen_ip: 127.0.0.1
    listen_port: ""
  dns:
    enabled: false
    listen_ip: 127.0.0.1
    domain: ""
  client:
    http_base_url: ""
    dns_server_ip: ""
    remote_server: false
```
 - store_events server 是否存储 event，用于 debug
 - token 是用于防止 api 被非法调用
 - http
   - listen_ip 监听的 ip
   - listen_port 如果是空字符串，就自动选择一个
 - dns （如果启用，可能需要权限才能监听 53 端口）
   - domain 在 dns 查询的时候的一级域名。如果此域名的 ns 服务器就是反连平台的地址，那么直接使用 `dig random.domain.com` 就可以让 dns 请求到反连平台，否则需要 `dig random.domain.com @reverse-server-ip` 指定 dns 服务器才可以。如果为空，则使用 `example.com`，dig 的时候需要指定 dns。
 - client
   - base_url 是客户端访问的时候使用，详见下方的场景
   - remote_server 是客户端访问的时候，是 http fetch 还是直接代码访问，详见下方的场景。非 remote_server 的时候，将自动生成，可以不填，否则必须两边一致。
   - dns_server_ip 是客户端发起 dns 查询的时候，使用的 ip，详见下方场景。

## 场景分析

### 场景1 - 扫描器和靶站可以使用 ip 双向互联

可以使用默认配置，配置监听 ip 就可以。比如 `your-reverse-server-ip` 为 `192.168.1.2`。

```yaml
reverse:
  store_events: false
  token: ""
  http:
    enabled: true
    listen_ip: your-reverse-server-ip
    listen_port: ""
  dns:
    enabled: true
    listen_ip: your-reverse-server-ip
    domain: ""
  client:
    http_base_url: ""
    dns_server_ip: ""
    remote_server: false
```

这样扫描器选择一个未占用的端口，生成 `base_url`，值为 `http://$http:listen_ip:$http:listen_port` 让被靶站尝试去访问。

DNS 相关测试中，就会使用 `dig some-domain @$dns.listen_ip` 的命令。

### 场景2 - 扫描器 listen 的地址和靶站访问的地址并不一样

适用于以下情况，需要指定 base_url

#### 一些云主机，虽然公网 ip 可以访问，但是本地并无法直接 listen 那个 ip

```yaml
reverse:
  store_events: false
  token: ""
  http:
    enabled: true
    listen_ip: 0.0.0.0
    listen_port: ""
  dns:
    enabled: false
    listen_ip: 0.0.0.0
    domain: ""
  client:
    # 上面的 port 留空代表自动选择，这里 ${port} 引用上面自动选择的值
    http_base_url: "http://your-reverse-server-ip:${port}"
    dns_server_ip: "your-reverse-server-ip"
    remote_server: false
```

#### 想使用解析到这个 ip 的域名让靶站访问

```yaml
reverse:
  store_events: false
  token: ""
  http:
    enabled: true
    listen_ip: 0.0.0.0
    listen_port: ""
  dns:
    enabled: false
    listen_ip: 0.0.0.0
    domain: ""
  client:
    # 上面的 port 留空代表自动选择，下面的 ${port} 引用上面自动选择的值
    http_base_url: "http://your-reverse-server-domain:${port}"
    // dns_server 只能是 ip
    dns_server_ip: "your-reverse-server-ip"
    remote_server: false
```

### 场景3 - 扫描器可以访问靶站，但是靶站无法访问扫描器。

这是非常常见的情况，比如在个人电脑上运行扫描器，扫描公网的靶站。这时候需要在公网上也部署一份反连平台，然后扫描器和靶站都去使用那个。

对于单独部署的反连平台，使用 `./xray reverse` 启动，配置如下。

```yaml
reverse:
  store_events: false
  token: "token-value"
  http:
    enabled: true
    listen_ip: 0.0.0.0
    listen_port: "12345"
  dns:
    enabled: true
    listen_ip: 0.0.0.0
    domain: ""
  client:
    http_base_url: ""
    dns_server_ip: ""
    remote_server: false
```

对于扫描器端，配置如下。

```yaml
reverse:
  store_events: false
  token: "token-value"
  http:
    enabled: false
    listen_ip: 127.0.0.1
    listen_port: ""
  dns:
    enabled: false
    listen_ip: 127.0.0.1
    domain: ""
  client:
    http_base_url: "http://your-reverse-server-ip-or-domain:12345"
    dns_server_ip: "your-reverse-server-ip"
    remote_server: true
```

要注意的是，两边的 xray 请使用相同版本的，否则可能存在 api 不兼容的问题。

## HTTP API

!> **注意** 本 api 在开发阶段可能经常修改，请尽量使用最新版本，如有问题可以反馈。

在 `remote_server` 模式下，扫描器和反连平台是通过 HTTP API 通信的，下面简单描述一下这些 API。

注意，以下的 api 都需要 token 和 groupID 配合使用，下文中引用的 hashedToken 非配置文件中的 token 值，而是 `Sha256(token + groupID + unitID)[:6]` 得到的值。

### 生成一个访问反连平台的 http url

此 url 的生成不需要通信，直接按照规则拼接即可

```go
fmt.Sprintf("%s/i/%s/%s/%s/", HTTPBaseURL, hashedToken, group.id, unit.id)
```
`HTTPBaseURL` 的生成规则见上文。

其中 `group` 和 `unit` 是比较重要的两个概念。一个 group 可以包括很多个 unit，都有自己的 id。设想一个场景，一个可能存在 ssrf 的输入点扫描器会构造很多种输入用于 fuzz，但是希望有一个输入造成了回连就可以认定有漏洞了，否则就会输出很多个重复的漏洞，这时候可以将这些归为一个 group，多个 unit。

在后文的 log 查询 api 中，都是使用 `group.id` 去查询，将返回这个 group 下面第一个 unit 的结果，然后删除这个 group 下面的其他结果。如果这个 group 有新的请求在查询之后到来，继续查询是可以查询到的。因为 `group.id` 和 `unit.id` 的生成并不需要向反连平台注册。

如果并不需要这样的去重，可以每次都生成一个新的 group，每个 group 下面只有一个 unit。

下文的样例都假设 token 为 `imtoken`，本样例中 `d6f7be` 的来源就是 `hashlib.sha256(b"imtoken" + b"a" + b"b").hexdigest()[:6]`。

```
url http://127.0.0.1:9999/i/d6f7be/a/b/ -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /i/d6f7be/a/b/ HTTP/1.1
> Host: 127.0.0.1:9999
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 07 Aug 2019 03:21:04 GMT
< Content-Length: 22
<
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"data":null}
```

要注意的是这个 url 后面是可以任意追加的，只要保持前缀不变即可。比如 `http://127.0.0.1:9999/i/d6f7be/a/b/index.php?foo=bar`。

### 生成 dns log 的域名

```go
fmt.Sprintf("i-%s-%s-%s.%s", hashedToken, group.id, unit.id, Domain)
```

`group` 和 `unit` 的含义和上文一致，`Domain` 的含义是根域名，详见上文配置文件中相关的部分。注意，`id` 需要符合域名的规则，建议只有小写字母和数字，否则解析可能会出错。

目前反连平台支持 A 和 AAAA 记录，解析结果均为 `127.0.0.1` 或者 `::1`。

```
dig i-d6f7be-a-b.example.com A @127.0.0.1

; <<>> DiG 9.10.6 <<>> i-d6f7be-a-b.example.com A @127.0.0.1
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 37454
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 0
;; WARNING: recursion requested but not available

;; QUESTION SECTION:
;i-d6f7be-a-b.example.com.	IN	A

;; ANSWER SECTION:
i-d6f7be-a-b.example.com. 60	IN	A	127.0.0.1

;; Query time: 0 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Wed Aug 07 11:24:30 CST 2019
;; MSG SIZE  rcvd: 82
```

### 查询 http / dns log

api url 为 `/fetch/{token}/:group`

`{token}` 代表配置文件中的 token

`:group` 代表这是一个来自 url 中的变量，取值是 `group.id`。

返回值是

```go
const (
	EventTypeHTTPVisit eventType = 0
	EventTypeDNSQuery  eventType = 1
)

type Event struct {
	UnitId    string
	TimeStamp int64
	EventType eventType
	// 字符串，方便去序列化 http 传输等
	// http 下就是原始请求，dns 下就是请求的域名
	Request    string
	RemoteAddr string
}
```

http log 

```
curl http://127.0.0.1:9999/fetch/imtoken/a -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /fetch/imtoken/a HTTP/1.1
> Host: 127.0.0.1:9999
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 07 Aug 2019 03:26:17 GMT
< Content-Length: 145
<
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"data":{"unit_id":"b","time_stamp":1565148270333,"event_type":1,"request":"i-d6f7be-a-b.example.com.","remote_addr":"127.0.0.1:61277"}}```

dns log

```
curl http://127.0.0.1:9999/fetch/imtoken/a -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /fetch/imtoken/a HTTP/1.1
> Host: 127.0.0.1:9999
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 07 Aug 2019 03:27:35 GMT
< Content-Length: 145
<
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"data":{"unit_id":"b","time_stamp":1565148452139,"event_type":1,"request":"i-d6f7be-a-b.example.com.","remote_addr":"127.0.0.1:61265"}}
```

如果查询不到结果，将返回 null

```
curl http://127.0.0.1:9999/fetch/imtoken/notexist -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /fetch/imtoken/notexist HTTP/1.1
> Host: 127.0.0.1:9999
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 07 Aug 2019 03:27:50 GMT
< Content-Length: 22
<
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"data":null}
```

### debug 查看所有的 log

如果在配置文件中开启了 `store_events: true`，反连平台会将收到的请求都记录下来，方便人工查看和调试。注意，本功能不影响上述的 api 的行为，该结果没有持久化，重启会丢失。

浏览器访问 `/list_events/{token}`

```
UnitID: b Timestamp: 1564624330660 IP: 127.0.0.1:63527

GET /v/z92dai/a/b HTTP/1.1
Host: 127.0.0.1:9999
Accept: */*
User-Agent: curl/7.54.0

UnitID: y Timestamp: 1564624376515 IP: 127.0.0.1:54789

z92dai-x-y.example.com.
```

### 健康检查

如果想知道反连平台是否启动，可以访问 `/health_check/{token}`。

```
curl http://127.0.0.1:9999/health_check/imtoken -v
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 9999 (#0)
> GET /health_check/imtoken HTTP/1.1
> Host: 127.0.0.1:9999
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 07 Aug 2019 03:29:16 GMT
< Content-Length: 22
<
* Connection #0 to host 127.0.0.1 left intact
{"code":0,"data":null}
```

