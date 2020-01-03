!> 注意，此功能只在高级版中提供

```yaml
subdomain:
  modes: # 使用哪些方式获取子域名
    - brute # 字典爆破模式
    - api # 使用各大 api 获取
    - zone_transfer # 尝试使用域传送漏洞获取
  worker_count: 100 # 决定同时允许多少个 DNS 查询
  dns_servers: # 查询使用的 DNS server
    - 1.1.1.1
    - 8.8.8.8
  allow_recursive: false # 是否允许递归扫描，开了后如果发现 a.example.com 将继续扫描 a.example.com 的子域名
  max_depth: 5 # 最大允许的子域名深度
  main_dictionary: "" # 一级子域名字典， 绝对路径
  sub_dictionary: "" # 其它层级子域名字典， 绝对路径
```

子域名的配置项相对比较简洁，对照注释大都可以理解。

## worker_count

这个 worker_count 对应于 goroutine, 如果没写过 go，可以理解为其他语言的协程、线程。如果设置为 100，指的是同时可能有 100 个 DNS 查询请求发出。这和 http 配置中的 `max_qps` 不一样，`max_qps` 指的是 1s 内最大运行的请求数。

## `allow_recursive` 和 `max_depth`

假设子域名扫描的目标为  `example.com`, 发现有子域名 `a.example.com`。
当开启 `allow_recurisive` 后，将自动把 `a.example.com` 视为新的扫描目标，进而可能获得 `b.a.example.com` 等二级子域。 而最大的子域深度由 `max_depth` 控制。

##  `main_dictionary` 和 `sub_dictionary`

与 `dirscan` 插件类似，当没有配置这两项时将使用内置字典，默认 main 字典为 Top3000, 默认 sub 字典为 top200。
当配置了自定义字典时，将使用用户配置的字典而禁用内直字典，两个配置项可以单独配置。