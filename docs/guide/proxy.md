# MITM 代理扫描

## MITM 代理会做什么

配置文件中有一段配置如下

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

目前代理发送请求只会受到配置文件中 `http` 部分的部分限制，比如 HTTP 头、cookie、proxy 等是不影响代理的，如果想串联代理，需要使用 pupstream_proxy 字段。

## 配合 burp 使用

### burp 作为 xray 的下级代理

场景：xray 作为浏览器的代理，burp 作为 xray 的代理

场景：
 - xray 做漏洞扫描
 - burp 查看和记录 xray 的扫描记录和原始请求
 
查看 burp [Proxy -> Options -> Porxy Listeners]，作为 xray 配置文件中的 [mitm -> upstream_proxy] 的值。
 
### xray 作为 burp 的下级代理

burp 作为浏览器的代理，xray 作为 burp 的代理

场景：
 - burp 查看和记录浏览器的请求
 - xray 做漏洞扫描
 
查看 burp [Project Options -> Connections -> Upstream Proxy Servers]，勾选 `Override user options`，点击 `Add`，将 xray 的 mitm 代理配置填写在里面。

比如 `./xray webscan --listen 127.0.0.1:8072`，那么 `Proxy Host` 和 `Proxy port` 就是 `127.0.0.1` 和 `8072`。`Authentication type` 如果在 xray 配置文件中设置了用户名密码，就选择 `Basic`，然后填写用户名密码。
