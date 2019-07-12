# 反连平台

## 什么是 ssrf，反连平台是干什么的

![](https://chaitin.github.io/xray/assets/reverse.jpg)

## 反连平台配置项目

```yaml
reverse:
  listen_ip: 127.0.0.1
  listen_port: ""
  store_events: false
  base_url: ""
  remote_server: false
```

 - listen_ip 监听的 ip
 - listen_port 如果是空字符串，就自动选择一个
 - store_events server 是否存储 event，用于 debug
 - base_url 是客户端访问的时候使用，详见下方的场景
 - remote_server 是客户端访问的时候，是 http fetch 还是直接代码访问，详见下方的场景
 - token 是用于防止 api 被非法调用

### 场景1 - 扫描器和靶站可以使用 ip 双向互联

可以使用默认配置，配置监听 ip 就可以。

```yaml
reverse:
  listen_ip: some-ip
  # 自动选择端口
  listen_port: ""
  store_events: false
  base_url: ""
  remote_server: false
```

这样扫描器就会生成 `base_url`，值为 `http://$listen_ip:$listen_port` 让被靶站尝试去访问。

## 场景2 - 扫描器 listen 的地址和靶站访问的地址并不一样

适用于以下情况，需要指定 base_url

### 一些云主机，虽然公网 ip 可以访问，但是本地并无法直接 listen 那个 ip

```yaml
reverse:
  listen_ip: 0.0.0.0
  listen_port: ""
  store_events: false
  # 上面的 port 留空代表自动选择，下面的 ${port} 引用上面自动选择的值
  base_url: "http://some-ip:${port}"
  remote_server: false
```

### 想使用解析到这个 ip 的域名让靶站访问

```yaml
reverse:
  listen_ip: 0.0.0.0
  listen_port: ""
  store_events: false
  base_url: "http://some-domain:${port}"
  remote_server: false
```

### 扫描器反连平台前面有端口映射
 
 ```yaml
reverse:
  listen_ip: 0.0.0.0
  listen_port: "some_port"
  store_events: false
  base_url: "http://some-ip:real_port"
  remote_server: false
```

### 场景3 - 扫描器可以访问靶站，但是靶站无法访问扫描器。

这是非常常见的情况，比如在个人电脑上运行扫描器，扫描公网的靶站。这时候需要在公网上也部署一份反连平台，然后扫描器和靶站都去使用那个。

对于单独部署的反连平台，使用 `./xray reverse` 启动，配置如下。

```yaml
reverse:
  listen_ip: "0.0.0.0"
  listen_port: "port"
  token: "your-token"
```

对于扫描器端，配置如下。

```yaml
reverse:
  base_url: "http://my-reverse-server:port"
  remote_server: true
  token: "your-token"
```
