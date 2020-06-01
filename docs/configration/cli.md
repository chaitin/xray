# 命令详解

> 这里以 mac 平台的二进制作为演示，所以二进制为 `xray_darwin_amd64`，其他平台请自行更改。

```
$ ./xray_darwin_amd64 -h

 __   __  _____              __     __
 \ \ / / |  __ \      /\     \ \   / /
  \ V /  | |__) |    /  \     \ \_/ /
   > <   |  _  /    / /\ \     \   /
  / . \  | | \ \   / ____ \     | |
 /_/ \_\ |_|  \_\ /_/    \_\    |_|


Version: 0.14.0/d1742479/COMMUNITY

NAME:
   xray - A powerful scanner engine [https://www.chaitin.cn/zh/xray]

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
     webscan    Run a webscan task
     reverse    Run a standalone reverse server
     genca      Generate CA certificate and key
     subdomain  Run a subdomain task
     version    Show version info
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE      Load configuration from FILE
   --log_level value  Log level, choices are debug, info, warn, error, fatal
   --help, -h         show help
```

## GLOBAL OPTIONS

先来看一个界面中的这三个全局配置项。全局配置的意思是如果在这指定了，那么所有命令执行的时候都会生效。

+ `--config` 用于指定配置文件的位置，默认加载同目录的 `config.yaml`。关于配置文件请看下一节文档的内容。
+ `--log_level` 用于指定全局的日志配置，默认为 `info`， 可以通过设置为 `debug` 来查看更详细的请求信息、运行时日志信息。

全局配置的使用时需要紧跟二进制程序，如:

```
./xray_darwin_amd64 --log_level debug --config 1.yaml webscan --url xxx
./xray_darwin_amd64 --log_level debug reverse
```

而下面这种方式是不生效的，使用时需要注意下:

```
./xray_darwin_amd64 webscan --log_level debug
```

## COMMANDS

xray 的命令有 6 个，抛开 `version` 和 `help` 这两个信息展示型的命令，还有 `webscan`, `reverse`, `genca`, `subdomain` 四个。

`reverse` 命令用于启动单独的盲打平台服务，盲打平台用于处理没有回显或延迟触发的问题，如果你挖过存储型 XSS，一定对这个不陌生。这部分内容相对独立，单独用一节来介绍。

`genca` 用于快速生成一个根证书，主要用于被动代理扫描 HTTPS 流量时用到。

`subdomain` 是子域名扫描的命令，仅**高级版**才有。

`webscan` 是 xray 的重头戏。

这里介绍一下后面两个命令。

## subdomain 子域名扫描

扫描 `example.com`，并将结果输出到 example.txt

```
./xray_darwin_amd64  subdomain --target example.com --text-output example.txt
```

扫描 `example.com`,并使用 console ui 交互式界面，同时记录结果到 example.txt

```
./xray_darwin_amd64  subdomain --target example.com --console-ui --text-output example.txt
```
![cui.svg](../assets/configuration/cui.svg)

其他用法请参照 subdomain 配置文件中的内容

## webscan web 漏洞检测
运行 `./xray_darwin_amd64 webscan -h`，可以看到

```
NAME:
    webscan - Run a webscan task

USAGE:
    webscan [command options] [arguments...]

OPTIONS:
   --plugins value         specify the plugins to run, separated by ','
   --poc value             specify the poc to run, separated by ','

   --listen value          use proxy resource collector, value is proxy addr
   --basic-crawler value   use a basic spider to crawl the target and scan the results
   --url-file value        read urls from a local file and scan these urls
   --url value             scan a **single** url
   --data value            data string to be sent through POST (e.g. 'username=admin')
   --raw-request FILE      load http raw request from a FILE

   --json-output FILE      output xray results to FILE in json format
   --html-output FILE      output xray result to FILE in HTML format
   --webhook-output value  post xray result to url in json format
```

为了方便大家理解，这里其实将 OPTIONS 分为了三部分:

#### 配置扫描插件

`--plugins` 配置本次扫描启用哪些插件, 不再使用配置文件中的配置

```
--plugins xss
--plugins xss,sqldet,phantasm
```

`--poc` 配置本次扫描启用哪些 POC, 因为所有 POC 隶属于 phantasm 插件, 所以该参数其实是 phantasm 插件独有的配置。为了使用方便，该参数支持 Glob 表达式批量加载，解析规则为用该参数值匹配内置的 poc 名字，如果有匹配到则启用；然后检查能否匹配本地文件，如果能匹配到，也加载。用起来是非常灵活的，如：

```
只加载一个 POC, 精准匹配
--plugins phantasm --poc poc-yaml-thinkphp5-controller-rce

加载内置的所有带 `thinkphp` 的 POC
--plugins phantasm --poc "*thinkphp*"

加载本地 `/home/test/pocs/` 目录所有的 POC:
--plugins phantasm --poc "/home/test/pocs/*"

加载 `/home/test/pocs/` 下包含 thinkphp 的 POC
--plugins phantasm --poc "/home/test/pocs/*thinkphp*"

...
```

#### 配置输入来源

中间的这四个是互斥的，意味着一次只能启用这5个的一个。

`--listen` 启动一个被动代理服务器作为输入，如 `--listen 127.0.0.1:4444`，然后配置浏览器或其他访问工具的 http 代理为 http://127.0.0.1:4444 就可以自动检测代理中的 HTTP 请求并进行漏洞扫描。

`--basic-crawler` 启用一个基础爬虫作为输入， 如 `--basic-crawler http://example.com`，就可抓取 http://example.com 的内容并以此内容进行漏洞扫描

`--url` 用于快速测试单个 url，这个参数不带爬虫，只对当前链接进行测试。默认为 GET 请求，配合下面的 `--data` 参数可以指定 body，同时变为 POST 请求。
```
--url http://example.com/?p=q --data "x=y&a=b"

将产生如下的请求进行探测:

POST http://example.com/?p=q HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169
Content-Length: 7
Content-Type: application/x-www-form-urlencoded
Accept-Encoding: gzip

a=b&x=y
```

`--raw-request` 用于加载一个原始的 HTTP 请求并用于扫描，原始请求类似上面代码框中的原始请求，如果你用过 `sqlmap -r`，那么这个参数应该也很容易上手。

#### 配置输出

最后三个用于指定结果输出方式，这三种方式可以单独使用，也可以搭配使用。

+ `--html-output` 将结果输出为 html 报告, [报告样例](../assets/report_example.html)
+ `--webhook-output` 将结果发送到一个地址
+ `--json-output` 将结果输出到一个 json 文件中

`--webhook-output`和`--json-output` 输出是 json 格式的结构化数据，数据格式参照: [漏洞格式](api/vuln.md)。

你可以在`--json-output`和`--html-otput`参数中使用变量`__timestamp__`和` __datetime__`，这样文件名中对应位置会自动替换为时间戳或日期时间，避免输出到同一文件时报错。如`--html-output report-__datetime__.html`将使用`report-2019_11_01-10_03_26.html`作为报告文件名。


## 联合使用

将上面说的一些结合起来使用，就可以满足多种场景下的使用需求了。下面的例子都是可以正常运行的，作用不言而喻。

```
./xray_darwin_amd64 webscan --plugins xss --listen 127.0.0.1:1111 --html-output 1.html

./xray_darwin_amd64 --log_level debug webscan --plugins xss,cmd_injection --basic-crawler http://example.com --json-output 1.json

./xray_darwin_amd64 webscan --url http://example.com --data "x=y" --html-output 2.html --json-output 1.json

./xray_darwin_amd64 webscan --url http://example.com/ --webhook-output http://host:port/path
```

## 交互式命令行

如果你已经理解了上面的这些内容，但感觉内容繁多记不住，那这个交互式的命令行就是新手福利了。

直接运行 xray 而不加任何参数即可启动交互式命令行。

![ui](../assets/configuration/terminalui.svg)
