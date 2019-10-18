# 配置文件

引擎初次运行时，会在当前目录内生成一个 `config.yaml` 文件，该文件中的配置项可以直接左右引擎在运行时的状态。通过调整配置中的各种参数，可以满足不同场景下的需求。在修改某项配置时，请务必理解该项的含义后再修改，否则可能会导致非预期的情况。

> 在 xray 快速迭代时期，不保证配置文件向后兼容。如果出错，可以备份配置文件并重新生成。
> 实际上建议每次更新版本后都备份配置文件后删除并重新生成，以免错过新功能的配置。

在这里给一份 xray 中默认的完整配置项，方便快速查阅。具体细节及说明请点击相关的小章节。

```yaml

version: 2.1
plugins:
  xss:
    enabled: true
    ie_feature: false
  baseline:
    enabled: true
    detect_outdated_ssl_version: false
    detect_http_header_config: false
    detect_cors_header_config: true
    detect_server_error_page: false
    detect_china_id_card_number: false
    detect_serialization_data_in_params: true
  cmd_injection:
    enabled: true
    detect_blind_injection: false
  crlf_injection:
    enabled: true
  dirscan:
    enabled: true
    dictionary: ""
  jsonp:
    enabled: true
  path_traversal:
    enabled: true
  redirect:
    enabled: true
  sqldet:
    enabled: true
    error_based_detection: true
    boolean_based_detection: true
    time_based_detection: true
    # 下面两个选项很危险，开启之后可以增加检测率，但是有破坏数据库数据的可能性，请务必了解工作原理之后再开启
    dangerously_use_comment_in_sql: false
    dangerously_use_or_in_sql: false
  ssrf:
    enabled: true
  xxe:
    enabled: true
  upload:
    enabled: true
  brute_force:
    enabled: true
    username_dictionary: ""
    password_dictionary: ""

  phantasm:
    enabled: true
    depth: 1
    poc:
      - poc-yaml-activemq-cve-2016-3088
      - poc-yaml-bash-cve-2014-6271
      - poc-yaml-cacti-weathermap-file-write
      - poc-yaml-coldfusion-cve-2010-2861-lfi
      - poc-yaml-confluence-cve-2015-8399
      - poc-yaml-confluence-cve-2019-3396-lfi
      - poc-yaml-coremail-cnvd-2019-16798
      - poc-yaml-couchdb-cve-2017-12635
      - poc-yaml-couchdb-unauth
      - poc-yaml-dedecms-url-redirection
      - poc-yaml-discuz-wooyun-2010-080723
      - poc-yaml-discuz-v72-sqli
      - poc-yaml-discuz-wechat-plugins-unauth
      - poc-yaml-dlink-850l-info-leak
      - poc-yaml-dlink-cve-2019-16920-rce
      - poc-yaml-docker-api-unauthorized-rce
      - poc-yaml-docker-registry-api-unauth
      - poc-yaml-drupal-drupalgeddon2-rce
      - poc-yaml-drupalgeddon-cve-2014-3704-sqli
      - poc-yaml-ecology-javabeanshell-rce
      - poc-yaml-ecology-workflowcentertreedata-sqli
      - poc-yaml-ecshop-360-rce
      - poc-yaml-elasticsearch-cve-2015-1427
      - poc-yaml-elasticsearch-unauth
      - poc-yaml-etouch-v2-sqli
      - poc-yaml-finereport-directory-traversal
      - poc-yaml-glassfish-cve-2010-2861-lfi
      - poc-yaml-hadoop-yarn-unauth
      - poc-yaml-ifw8-router-cve-2019-16313
      - poc-yaml-influxdb-unauth
      - poc-yaml-jboss-unauth
      - poc-yaml-jenkins-cve-2018-1000861-rce
      - poc-yaml-joomla-cve-2015-7297-sqli
      - poc-yaml-joomla-cve-2017-8917-sqli
      - poc-yaml-joomla-ext-zhbaidumap-cve-2018-6605-sqli
      - poc-yaml-maccmsv10-backdoor
      - poc-yaml-metinfo-cve-2019-16996-sqli
      - poc-yaml-metinfo-cve-2019-16997-sqli
      - poc-yaml-metinfo-cve-2019-17418-sqli
      - poc-yaml-metinfo-lfi-cnvd-2018-13393
      - poc-yaml-nextjs-cve-2017-16877
      - poc-yaml-nhttpd-cve-2019-16278
      - poc-yaml-nuuo-file-inclusion
      - poc-yaml-php-cgi-cve-2012-1823
      - poc-yaml-phpmyadmin-cve-2018-12613-file-inclusion
      - poc-yaml-phpmyadmin-setup-deserialization
      - poc-yaml-phpstudy-backdoor-rce
      - poc-yaml-phpunit-cve-2017-9841-rce
      - poc-yaml-joomla-cnvd-2019-34135-rce
      - poc-yaml-springcloud-cve-2019-3799
      - poc-yaml-tomcat-cve-2018-11759
      - poc-yaml-pulse-cve-2019-11510
      - poc-yaml-rails-cve-2018-3760-rce
      - poc-yaml-resin-cnnvd-200705-315
      - poc-yaml-resin-inputfile-fileread-or-ssrf
      - poc-yaml-resin-viewfile-fileread
      - poc-yaml-seacms-v654-rce
      - poc-yaml-apache-solr-cve-2017-12629-xxe
      - poc-yaml-solr-cve-2019-0193
      - poc-yaml-supervisord-cve-2017-11610
      - poc-yaml-tensorboard-unauth
      - poc-yaml-thinkphp5-controller-rce
      - poc-yaml-thinkphp5023-method-rce
      - poc-yaml-tomcat-cve-2017-12615-rce
      - poc-yaml-uwsgi-cve-2018-7490
      - poc-yaml-vbulletin-cve-2019-16759
      - poc-yaml-weblogic-cve-2017-10271-reverse
      - poc-yaml-weblogic-cve-2019-2729-1
      - poc-yaml-weblogic-cve-2019-2729-2
      - poc-yaml-weblogic-ssrf
      - poc-yaml-weblogic-cve-2017-10271
      - poc-yaml-weblogic-cve-2019-2725
      - poc-yaml-webmin-cve-2019-15107-rce
      - poc-yaml-wuzhicms-v410-sqli
      - poc-yaml-yungoucms-sqli
      - poc-yaml-zabbix-authentication-bypass
      - poc-yaml-zabbix-cve-2016-10134-sqli
      - poc-yaml-zimbra-cve-2019-9670-xxe
      - poc-go-tomcat-put

log:
  level: info # 支持 debug, info, warn, error, fatal

mitm:
  ca_cert: ./ca.crt
  ca_key: ./ca.key
  auth:
    username: ""
    password: ""
  restriction:
    includes: # 允许扫描的域
    - '*' # 表示允许所有的域名和 path
    - "example.com/admin*" # 表示允许 example.com 下的 /admin 开头的 path
    excludes:
    - '*google*'
    - '*github*'
    - '*.gov.cn'
    - '*.edu.cn'
  allow_ip_range: []
  queue:
    max_length: 10000
  proxy_header:
    via: "" # 如果不为空，proxy 将添加类似 Via: 1.1 $some-value-$random 的 http 头
    x_forwarded: false # 是否添加 X-Forwarded-{For,Host,Proto,Url} 四个 http 头
  upstream_proxy: "" # mitm 的全部流量继续使用 proxy

basic_crawler:
  max_depth: 0 # 爬虫最大深度, 0 为无限制
  max_count_of_links: 0 # 本次扫描总共爬取的最大连接数， 0 为无限制
  allow_visit_parent_path: false # 是否允许访问父目录, 如果扫描目标为 example.com/a/， 如果该项为 false, 那么就不会爬取 example.com/ 这级目录的内容
  restriction: # 和 mitm 中的写法一致, 有个点需要注意的是如果当前目标为 example.com 那么会自动添加 example.com 到 includes 中。
    includes: []
    excludes:
    - '*google*'

reverse:
  db_file_path: ""
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
    # 静态解析规则
    resolve:
    - type: A # A, AAAA, TXT 三种
      record: localhost
      value: 127.0.0.1
      ttl: 60
  client:
    http_base_url: ""
    dns_server_ip: ""
    remote_server: false
http:
  proxy: "" # 漏洞扫描时使用的代理
  dial_timeout: 5 # 建立 tcp 连接的超时时间
  read_timeout: 30 # 读取 http 响应的超时时间，不可太小，否则会影响到 sql 时间盲注的判断
  fail_retries: 1 # 请求失败的重试次数，0 则不重试
  max_redirect: 5 # 单个请求最大允许的跳转数
  max_qps: 500 # 每秒最大请求数
  max_conns_per_host: 50 # 同一 host 最大允许的连接数，可以根据目标主机性能适当增大。
  max_resp_body_size: 8388608 # 8M，单个请求最大允许的响应体大小，超过该值 body 就会被截断
  headers: # 每个请求预置的 http 头
    User-Agent:
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
    - PROPFIND
    - MOVE
  tls_skip_verify: true # 是否验证目标网站的 https 证书。
```
