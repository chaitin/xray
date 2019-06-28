# 反连平台

## 什么是 ssrf，反连平台是干什么的

![](https://github.com/chaitin/xray/blob/master/assets/reverse.jpg?raw=true)

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
 
### 场景1 - 本地搭建扫描器和 reverse server

而且本地有一个本地服务存在 ssrf，可以使用默认配置。

### 场景2 - 本地搭建扫描器和 reverse server，但是远程访问本地受限
 
比如 docker for mac 中运行了一个 ssrf 服务，访问 mac 上的 server，需要使用特定的域名 `docker.for.mac.localhost`

 ```yaml
reverse:
  listen_ip: 127.0.0.1
  listen_port: ""
  base_url: "http://docker.for.mac.localhost:${port}"
  remote_server: false
 ```
 
如果是访问本地必须使用域名或者反向代理或者端口映射，也是同理。

 ```yaml
reverse:
  listen_ip: 127.0.0.1
  listen_port: "8000"
  base_url: "http://some-domain"
  remote_server: false
  ```
 
### 场景3 - 本地搭建扫描器 但是远程无法访问本地
 
这种情况下，可以使用单独的 reverse server

对于 server 端，使用 `./xray reverse` 启动。
 
 ```yaml
reverse:
  listen_ip: 0.0.0.0
  listen_port: 8000
```
  
对于扫描器端
 
  ```yaml
 reverse:
   listen_ip: ""
   listen_port: ""
   base_url: "http://ip:8000"
   remote_server: true
  ```
  
这时候，ssrf 服务和扫描器都会去访问 `base_url`。
