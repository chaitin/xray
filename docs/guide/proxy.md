# HTTP 被动代理扫描
这一部分主要介绍配置项中 `mitm` 部分相关的内容。

```yaml
mitm:
  ca_cert: ./ca.crt
  ca_key: ./ca.key
  restriction:
    includes: # 允许扫描的域
    - '*' # 表示允许所有的域名和 path
    - "example.com/admin*" # 表示允许 example.com 下的 /admin 开头的 path
    excludes:
    - '*google*' 
  auth: # 代理使用密码保护
    username: ""
    password: ""
  proxy_header:
    via: "" # 如果不为空，proxy 将添加类似 Via: 1.1 $some-value-$random 的 http 头
    x_forwarded: false # 是否添加 X-Forwarded-{For,Host,Proto,Url} 四个 http 头
  upstream_proxy: "" # mitm 的全部流量继续使用 proxy
  queue:
    max_length: 10000
```

## 抓取 HTTPS 流量

和 burp 类似，抓取 https 流量需要信任一个根证书，这个根证书可以自行生成，也可用下列自带的命令生成:

```
xray genca
```

运行后将在当前目录生成 `ca.key` 和 `ca.crt`， 用户需要手动信任 `ca.crt`并按需调整配置中的文件位置，操作完成后就可以正常抓取 https 流量了。
> Firefox 浏览器没有使用系统的根证书管理器，意味着使用 Firefox 时需要单独在该浏览内导入证书才可生效。

## 限制漏洞扫描的范围
在 mitm 的配置中的 `restrction` 项指示本次漏洞的 URI 限制。

1. `includes`表示只扫描哪些域和路径。比如 `*.example.com` 只扫描 `example.com` 的子域
1. `excludes` 表示不扫描哪些域和路径。比如 `t.example.com` 表示不扫描 `t.example.com`

两个都配置的情况下会取交集，这两个配置常用于想要过滤代理中的某些域，或者只想扫描某个域的请求时。

两项配置都支持 path 过滤，如果输入的中有 `/`, 那么 `/` 后面的表达式就是 path 的过滤。可以对照上面 `yaml` 配置中的例子来理解。

## 代理启用密码保护

xray 支持给代理配置基础认证的密码，当设置好 `auth` 中的 `username` 和 `password` 后，使用代理时浏览器会弹框要求输出用户名密码，输入成功后代理才可正常使用。

## 队列配置
```yaml
  queue:
    max_length: 10000
```
在做被动扫描时，从被动代理中获取的请求往往不能一下被处理完，所以 xray 内部有个处理队列，表示当前堆积了多少请求没有被处理。xray 将每 10s 打印一下当前任务队列长度，一旦堆积的数量达到 `max_length`，代理将会 “卡住”，等待队列中的请求被处理后再继续生效。

## 请求头配置 `proxy_header`

```
proxy_header:
    via: "" # 如果不为空，proxy 将添加类似 Via: 1.1 $some-value-$random 的 http 头
    x_forwarded: false # 是否添加 X-Forwarded-{For,Host,Proto,Url} 四个 http 头
upstream_proxy: "" # mitm 的全部流量继续使用 proxy
```

如果开启 proxy_header，代理会添加 `via` 头和 `X-Forwarded-*` 系列头。如果在请求中就已经存在了同名的 HTTP 头，那么将会追加在后面。

比如 `curl http://127.0.0.1:1234 -H "Via: test" -H "X-Forwarded-For: 1.2.3.4" -v`，后端实际收到的请求将会是

```http
GET / HTTP/1.1
Host: 127.0.0.1:1234
User-Agent: curl/7.54.0
Accept: */*
Via: test, 1.1 xray-1fe7f9e5241b2b150f32
X-Forwarded-For: 1.2.3.4, 127.0.0.1
X-Forwarded-Host: 127.0.0.1:1234
X-Forwarded-Proto: http
X-Forwarded-Url: http://127.0.0.1:1234/
Accept-Encoding: gzip
```

## 代理的代理 `upstream_proxy`

!> 该项配置仅对 http 代理本身生效，不对漏洞扫描发出的请求生效。如果想配置漏洞扫描时的代理，请参照配置文件 `http` 部分的配置。

假如启动 xray 时配置的 listen 为 `127.0.0.1:7777`，设置该配置项为 `http://127.0.0.1:8080`, 那么浏览器设置代理为 `http://127.0.0.1:7777`，整体数据流如下:

此处有个图()

再次重申下，该配置仅影响代理本身，不会影响插件在漏洞探测时的发包行为，如果想代理漏洞探测的，请参照 `http` 的部分!