对于 web 扫描来说，http 协议的交互是整个过程检测过程的核心。因此这里的配置将影响到引擎进行 http 发包时的行为。

```yaml
http:
  proxy: "" # 漏洞探测使用的代理, 如 http://127.0.0.1:8080
  dial_timeout: 5 # 建立 tcp 连接的超时时间
  read_timeout: 30 # 读取 http 响应的超时时间，不可太小，否则会影响到 sql 时间盲注的判断
  fail_retries: 1 # 请求失败的重试次数，0 则不重试
  max_qps: 500 # 每秒最大请求数
  max_redirect: 5 # 单个请求最大允许的跳转数
  max_conns_per_host: 50 # 同一 host 最大允许的连接数，可以根据目标主机性能适当增大。
  max_resp_body_size: 8388608 # 8M，单个请求最大允许的响应体大小，超过该值 body 就会被截断
  headers: # 每个请求预置的 http 头
    UserAgent:
      - Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169
  cookies: # 每个请求预置的 cookie 值，效果上相当于添加了一个 Header: Cookie: key=value
    key: value
  allow_methods: # 允许使用 http 方法
    - HEAD
    - GET
    - POST
    - PUT
    - DELETE
    - OPTIONS
    - CONNECT
  tls_skip_verify: true # 是否验证目标网站的 https 证书。
```

## 漏洞扫描用的代理 `proxy`

配置该项后漏洞扫描发送请求时将使用代理发送，支持 `http`, `https` 和 `socks5` 三种格式，如:

```
http://127.0.0.1:1111
https://127.0.0.1:1111
socks5://127.0.0.1:1080
```

## 限制发包速度 `max_qps`

默认值 500， 因为最大允许每秒发送 500 个请求。一般来说这个值够快了，通常是为了避免被ban，会把该值改的小一些，极限情况支持设置为 1， 表示每秒只能发送一个请求。

## 其他配置项

对照配置文件的注释好好理解下应该就能懂，如果不懂就不要改了。
