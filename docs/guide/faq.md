## no command provided

在 Windows 下面直接双击打开或者使用了错误的参数，可能会提示 `No command provided, please run this program in terminal`，在其他平台下，会进入一个类似 shell 一样的界面。

这是因为 xray 是一款命令行工具，需要在命令行下运行并使用正确的参数，不能直接双击打开，具体的使用方法见 [下载运行](tutorial/prepare.md) 章节。

在非 Windows 环境下，没有参数运行将进入 shell 模式，可以继续输入 `version` 等指令运行，在 Windows 下面此特性被禁用。

## 网页上出现 xray 的报错信息 `Proxy Failed`

xray 在代理扫描的时候，需要同时和客户端到和服务端建立连接，如果 xray 和服务器端连接失败，就会返回这样一个错误页面给客户端，常见的错误解释如下

 - `timeout awaiting response` 等待代理返回时候超时，请检查网络情况，是否可以连通，如果是偶尔的超时，可以忽略，重试即可。
 - `connection reset by peer` 连接被中断，请检查对方是否有 waf，ip 是否被拉黑，如果是国外的目标，可能是网络不稳定的原因。
 - `certificate has expired or is not yet valid` 和 `certificate signed by unknown authority` 等 `x509` 开头的报错信息，请检查服务端的 ssl 证书配置是否正确，对于自签名证书等问题，可以选择不校验证书，将配置文件中的 `tls_skip_verify` 改为 `true` 即可。
 - `dial tcp: lookup xxx on xxx:53: no such host` 要访问的地址的域名，xray 解析失败。
   - 最常见原因是域名拼写错误，比如 `example.xray.cool` 写错为 `exmple.xray.cool`。
   - 在 MacOS 上，如果需要 VPN 才可以访问和解析的域名，在使用 xray 作为代理后，即使 VPN 连接了也无法解析。这是因为 xray 使用了跨平台编译技术，没有使用 MacOS 的 sdk，无法解析这种域名，普通域名不受影响。如果十分依赖这种场景，可以提交反馈。

## 使用了 `--json-output` 或者 `--html-output` 之后没有出现结果文件

xray 目前在扫到漏洞的时候才会创建文件并写入数据，否则说明没有发现漏洞。

## 如何确认代理模式下 xray 有没有在工作

代理模式下，如果 xray 接收到了流量并进行了扫描，每隔 10s 就会输出一条类似下方的消息

```
Statistic: scanned: 0, pending: 1, capacity: 10000, requestSent: 180, latency: 678.26ms, failedRatio: 0.00%
```

如果长期没有此信息输出，也没有漏洞等输出，说明 xray 可能没有接收到浏览器的流量，可以按照下面的办法排查

 - 访问 `http://proxy.xray.cool/` 或者 `http://xray/`，如果可以看到 `xray is working` 说明此浏览器已经成功设置 xray 为代理，请继续阅读下一条检查方法，否则请确认浏览器设置。
 - 浏览器等客户端代理可能存在白名单，比如一般 `127.0.0.1`、`localhost` 等默认不会走代理，这些白名单可能还存在用户设置的更多项目，请检查白名单。
 - 同理，xray 也存在扫描白名单，在配置文件的 `mitm` 部分中的 `restriction` 的 `includes` 和 `excludes` 中，请检查此部分配置，比如默认情况下 `gov` 和 `edu` 域名会被排除在外。

## 为什么使用 `--url` 指定目标扫描很快就结束了？

在指定 `--url URL` 的情况下，xray 只会针对该 URL 本身进行漏洞检测，不会爬取网站内容。比如，你输入 `--url https：//baidu.com`，xray 只会检测百度首页
的漏洞，这种情况下通常是无法扫描出漏洞的。

如果你需要利用爬虫爬取页面后再检测漏洞，可以使用 `--basic-crawler` 选项，具体使用方法请参考文档。

## 为什么我输入URL调用后扫不出漏洞，而其他扫描器扫出了很多漏洞？

原因同上，xray 社区版是一个被动扫描器（角色参照Burpsuite），主动爬虫功能较弱，可能无法满足你的需求。
建议将浏览器代理设置为xray的代理，通过代理的方式将网站流量发送给xray进行检测。如果还是不能扫出漏洞，可以反馈我们。

## 如何给xray扫描器设置上游代理？

xray 细化了代理的配置，大致分为两部分:

+ 如果要给xray的被动代理本身设置代理，修改配置文件 `mitm` 部分的 `upstream_proxy` 项
+ 如果要给xray的漏洞扫描设置代理可以在配置文件 http 部分增加 `proxy：http://127.0.0.1:7777` 这样的配置

## 如何限制xray发包速率？

在 `config.yaml` 文件中 `http` 设置 `max_qps`

## 为什么XSS输出的payload无法复现？

xray的xss插件检测方式与大家熟知的暴力发包方式不同，是通过分析语义来精准判断的，在此过程中可能不会发送真正的payload，而输出的payload也是通过分析得到的可能的payload，不保证可用。但一般来讲，不可用的情况也只需稍加调整就可以利用。

## xray会测请求包中的json数据吗？

会。 目前全局支持 query, url-encoded body, multipart body, json body，位置的漏洞探测。部分插件还会对 cookie 和 header 中的项进行漏洞探测。

## 如何批量爬取多个目标网站并进行漏洞扫描?

自带功能不支持。可以搭配第三方爬虫，同时使用 xray 的代理进行扫描。

## 代理模式下，为什么扫描完不退出

在代理模式下，xray 并不知道什么情况下是扫完了，因为 xray 是收到一个请求就扫描一个，比如你是浏览器挂 xray 为代理，没有收到新请求可能是你还没开始浏览页面，也可能是已经浏览完了关闭了浏览器，所以扫描完成之后需要手动的退出 xray 进程。
