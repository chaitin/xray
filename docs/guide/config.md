# 配置文件

引擎初次运行时，会在当前目录内生成一个 `config.yaml` 文件，该文件中的配置项可以直接左右引擎在运行时的状态。通过调整配置中的各种参数，可以满足不同场景下的需求。在修改某项配置时，请务必理解该项的含义后再修改，否则可能会导致非预期的情况。下列进阶使用的方法均与该配置文件相关。

## 抓取 https 流量

mitm 的配置项主要用于被动扫描模式下的代理的配置。

```yaml
mitm:
  ca_cert: ./ca.crt
  ca_key: ./ca.key
  includes:
    - "*"
  excludes:
    - "*google*"
```

配置项中的前两项： `ca_cert` 和 `ca_key` 用于指定中间人的根证书路径。和 burp 类似，抓取 https 流量需要信任一个根证书，这个根证书可以自行生成，也可用下列自带的命令生成:

```
xray genca
```


运行后将在当前目录生成 `ca.key` 和 `ca.crt`， 用户需要手动信任 `ca.crt`。操作完成后就可以正常抓取 https 流量了。

在 mitm 的配置部分还有两项配置值得注意：

1. `includes`表示只扫描哪些域。比如 `*.example.com` 只扫描 `example.com` 的子域
1. `excludes` 表示不扫描哪些域。比如 `t.example.com` 表示不扫描 `t.example.com`

两个都配置的情况下会取交集，这两个配置常用于想要过滤代理中的某些域，或者只想扫描某个域的请求时。默认配置为抓取所有的域。

## HTTP 配置

这里的配置主要影响到引擎的 http 发包，如有需求，参考 yaml 中注释进行对应的修改。

```yaml
http:
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