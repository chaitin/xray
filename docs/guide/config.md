# 配置文件

引擎初次运行时，会在当前目录内生成一个 `config.yaml` 文件，该文件中的配置项可以直接左右引擎在运行时的状态。通过调整配置中的各种参数，可以满足不同场景下的需求。在修改某项配置时，请务必理解该项的含义后再修改，否则可能会导致非预期的情况。下列进阶使用的方法均与该配置文件相关。

> 在 xray 快速迭代时期，不保证配置文件向后兼容。如果出错，可以备份配置文件并重新生成。

## 被动扫描配置

xray 使用中间人（MITM) 的方式来获取代理中的请求，该部分内容较多，请前往该链接查看详情: [MITM 代理](guide/proxy.md)

## 基础爬虫配置

基础爬虫的配置项对应于 `basic-crawler` 部分，默认的配置如下，用法参照文件中的注释:

```yaml
basic_crawler:
  max_depth: 0 # 爬虫最大深度, 0 为无限制
  max_count_of_links: 0 # 本次扫描总共爬取的最大连接数， 0 为无限制
  allow_visit_parent_path: false # 是否允许访问父目录, 如果扫描目标为 example.com/a/，若该项为 false, 那么就不会爬取 example.com/ 这级目录的内容
  restriction: # 和 mitm 中的写法一致, 有个点需要注意的是如果当前目标为 example.com 那么会自动添加 example.com 到 includes 中。
    includes: []
    excludes:
    - '*google*'
```

## HTTP 配置

这里的配置主要影响到引擎进行 http 发包时的行为，如有需求，参考 yaml 中注释进行对应的修改。

```yaml
http:
  proxy: "" # 漏洞探测使用的代理, 也可以通过环境变量指定，如 export http_proxy=http://127.0.0.1:8080
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
