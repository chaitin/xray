## no command provided

在 Windows 下面直接双击打开或者使用了错误的参数，可能会提示 `No command provided, please run this program in terminal`，在其他平台下，会进入一个类似 shell 一样的界面。

这是因为 xray 是一款命令行工具，需要在命令行下运行并使用正确的参数，不能直接双击打开，具体的使用方法见 [下载运行](tutorial/prepare.md) 章节。

在非 Windows 环境下，没有参数运行将进入 shell 模式，可以继续输入 `version` 等指令运行，在 Windows 下面此特性被禁用。

## 网页上出现 xray 的报错信息 `Proxy Failed`

xray 在代理扫描的时候，需要同时和客户端到和服务端建立连接，如果 xray 和服务器端连接失败，就会返回这样一个错误页面给客户端，常见的错误解释如下

 - `timeout awaiting response` 等待代理返回时候超时，请检查网络情况，是否可以连通。
 - `connection reset by peer` 连接被中断，请检查对方是否有 waf，ip 是否被拉黑。
 - `certificate has expired or is not yet valid` 和 `certificate signed by unknown authority` 等 `x509` 开头的报错信息，请检查服务端的 ssl 证书配置是否正确，对于自签名证书等问题，可以选择不校验证书，将配置文件中的 `tls_skip_verify` 改为 `true` 即可。