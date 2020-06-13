!> 注意，此功能只在高级版中提供

```yaml
subdomain:
  max_parallel: 50 # 并发的 worker 数，类似线程数
  allow_recursion: false # 是否允许递归的处理子域名，开启后，扫描完一级域名后，会自动将一级的每个域名作为新的目标，去找二级域名, 递归层数由下面的配置决定
  max_recursion_depth: 3 # 最大允许的子域名层数，3 意为 3 级子域名
  web_only: false # 结果中仅显示有 web 应用的, 没有 web 应用的将被丢弃
  ip_only: false # 结果中仅展示解析出 IP 的，没有解析成功的将被丢弃
  servers: ["8.8.8.8", "8.8.4.4", "223.5.5.5", "223.6.6.6", "4.2.2.1", "114.114.114.114"] # 子域名扫描过程中使用的 DNS Server
  sources:
      brute: # 字典爆破模式, 会自动处理泛解析问题
          enabled: true
          main_dict: "" # 一级大字典路径，为空将使用内置的 TOP 30000 字典
          sub_dict: "" # 其他级小字典路径，为空将使用内置过的 TOP 100 字典
      httpfinder: # http 的一些方式来抓取子域名，包括 js, 配置文件，http header 等等
          enabled: true
      dnsfinder: # 使用 dns 的一些错误配置来找寻子域名，如区域传送（zone transfer)
          enabled: true
      certspotter: # 下面的都是 API 类的了
          enabled: true
      crt:
          enabled: true
      hackertarget:
          enabled: true
      qianxun:
          enabled: true
      rapiddns:
          enabled: true
      sublist3r:
          enabled: true
      threatminer:
          enabled: true
      virusTotal:
          enabled: true
```

子域名的配置项相对比较简洁，对照注释大都可以理解。