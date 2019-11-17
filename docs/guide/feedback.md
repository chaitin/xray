# 如何来反馈 xray 的问题

首先为了方便定位问题，我们都需要你运行的 xray 的版本。可以使用 `./xray version` `xray.exe version` 等。

样例的输出是

```
[xray 0.5.1/8f007251]
Build: [Fri Jul 12 14:10:46 CST 2019] [darwin/amd64] [RELEASE/COMMUNITY]
Compiler Version: go version go1.12.1 darwin/amd64
```

然后如果有版本更新，先尝试用新版本看看问题是否依然存在。

## 如果是误报

最好可以提供 `--json-output` 或者 `--html-output` 的报告，这样就不需要有太多的猜测和口头的信息传递。

提供之前请注意报告中可能含有敏感信息，可以复制特定的一部分。

如果没有报告，最好能提供复现的地址、插件名、打印出来的 payload以及你为什么认为这是误报等等。

## 如果是漏报

请提供下可以复现的地址，我们将尽快检查和完善。

## 如果是其他的问题

比如想提出需求，反馈进程 crash 等等，也需要尽可能多的细节。

# 反馈渠道

如有问题可以在 GitHub 提 issue, 也可在下方的讨论组里

1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ 群: 717365081
1. 微信群: 扫描以下二维码加我的个人微信，会把大家拉到 `xray` 官方微信群    

<img src="../assets/wechat.jpg" height="150px">
