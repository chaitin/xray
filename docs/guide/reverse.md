# 反连平台

## 什么是 ssrf，反连平台是干什么的

![](https://chaitin.github.io/xray/assets/reverse.jpg)

## 反连平台配置项目

```yaml
reverse:
  token: ""
  http:
    enabled: true
    listen_ip: 127.0.0.1
    listen_port: ""
  dns:
    enabled: false
    listen_ip: 127.0.0.1
    domain: ""
    is_domain_name_server: false
    resolve:
    - type: A
      record: localhost
      value: 127.0.0.1
      ttl: 60
  client:
    http_base_url: ""
    dns_server_ip: ""
    remote_server: false
```
 - token 是用于防止 api 被非法调用
 - http
   - listen_ip 监听的 ip
   - listen_port 如果是空字符串，就自动选择一个
 - dns （如果启用，可能需要权限才能监听 53 端口）
   - domain 在 dns 查询的时候的一级域名。如果此域名的 ns 服务器就是反连平台的地址，那么直接使用 `dig random.domain.com` 就可以让 dns 请求到反连平台，否则需要 `dig random.domain.com @reverse-server-ip` 指定 dns 服务器才可以。如果为空，则使用 `example.com`，dig 的时候需要指定 dns。`is_domain_name_server` 是指有没有配置 ns 服务器为反连平台的地址，用于提示扫描器内部 payload 的选择。
   - resolve 的配置类似常见的 dns 配置，如果反连平台收到配置的域名的解析请求，将按照配置的结果直接返回。
 - client
   - base_url 是客户端访问的时候使用，详见下方的场景
   - remote_server 是客户端访问的时候，是 http fetch 还是直接代码访问，详见下方的场景。非 remote_server 的时候，将自动生成，可以不填，否则必须两边一致。
   - dns_server_ip 是客户端发起 dns 查询的时候，使用的 ip，详见下方场景。

## 场景分析

### 场景1 - 扫描器和靶站可以使用 ip 双向互联

可以使用默认配置，配置监听 ip 就可以。比如 `your-reverse-server-ip` 为 `192.168.1.2`。

```yaml
reverse:
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

### 场景2 - 扫描器 listen 的地址和靶站访问的地址并不一样

适用于以下情况，需要指定 base_url

#### 一些云主机，虽然公网 ip 可以访问，但是本地并无法直接 listen 那个 ip

```yaml
reverse:
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

## 管理界面

新版的反连平台新增了管理界面，可以访问反连平台 http 地址，url 为 `/cland/`。