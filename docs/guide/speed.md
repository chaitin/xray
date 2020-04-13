# 优化扫描速度

## 评估扫描速度

我们要先评估一下扫描一个 url 需要发出多少请求，这个测试比较简单。

测试环境

 - 双核 8G 虚拟机
 - Ubuntu 18.04
 - loopback 网络

编写测试脚本，使用 xary 作为代理发送 `/?$randKey=a&b=2` 的请求，请注意第一个参数名是随机的，这样 xray 将不会对请求去重，模拟一个稍微极端的环境。

!> `/?$randKey=a&b=2` 为下文的 url 概念，真实环境下，参数数量可能会更多，参数数量是影响扫描请求量的重要因素

首先使用默认配置，开启了除了 `phantasm` （poc 模块）以外的所有模块，选择四条连续的统计数据（每条数据之间间隔为 5 秒钟）

一瞬间 pending 队列就满了

```
Statistic: scanned: 32, pending: 3000, capacity: 3000, requestSent: 4020, latency: 1.53ms, failedRatio: 0.00%
Statistic: scanned: 52, pending: 3000, capacity: 3000, requestSent: 6502, latency: 1.34ms, failedRatio: 0.00%
Statistic: scanned: 84, pending: 3000, capacity: 3000, requestSent: 9001, latency: 1.26ms, failedRatio: 0.00%
Statistic: scanned: 96, pending: 3000, capacity: 3000, requestSent: 11510, latency: 1.21ms, failedRatio: 0.00%
```

由上面的数据可以得到以下结果

 - 每秒扫描约 5 个 url
 - 每秒约发送 500 个请求
 - 一个 url 要发送约 90 个请求来扫描
 - cpu 占用率约 30%

按照这个数据，扫描 3000 个 url，需要约十分钟的时间，但是对于爬虫或者代理来说，生产 3000 个 url 是非常快的，这样就产生了不平衡，生产快，消费慢，`pending` 的数字就是已经生产还没消费的数量。

## 为什么会发送这么多请求

xray 有很多检测模块，如果模块发送的第一个 payload 就命中了漏洞，那就可以结束了，但是如果没命中漏洞，就需要更多的 payload 来扫描。

还比如说 `dirscan` 模块，内置了一个敏感文件规则库，为了能全面的发现敏感文件，势必会发送非常多的请求。

实际情况下，发送的请求量是上面的理论值乘以参数数量。

xray 的开发者会精心选择 payload 和设计扫描原理来降低请求量，比如

 - 参数类型判断，比如 `id=a` 可能就不需要发送 `id` 为数字的情况下的 payload
 - 场景判断，参数值满足特定情况才会进行扫描
 - 参数去重，比如 `id=1` `id=2` 不会重复扫描
 - 路径去重和限制深度，这些在配置文件中都可以自定义，比如 `/admin/backup.zip` 出现的概率比一个非常深的路径 `/example/foo/bar/1/2/3/backup.zip` 下面出现的概率要高很多。

## xray 的配置

在 xray 的默认配置中，影响扫描速度的主要有两个参数

 - `max_qps` 每秒最大请求数，默认值为 500
 - `max_parallel` 插件调度并发数，默认值为 30

因为我测试是内网环境，网络质量非常好，所以是明显 `max_qps` 限制了 xray 的扫描速度，将它调整为 2000。

```
Statistic: scanned: 160, pending: 3000, capacity: 3000, requestSent: 15563, latency: 3.29ms, failedRatio: 0.00%
Statistic: scanned: 256, pending: 3000, capacity: 3000, requestSent: 25357, latency: 2.84ms, failedRatio: 0.00%
Statistic: scanned: 372, pending: 3000, capacity: 3000, requestSent: 35106, latency: 2.53ms, failedRatio: 0.00%
Statistic: scanned: 480, pending: 3000, capacity: 3000, requestSent: 44902, latency: 2.40ms, failedRatio: 0.00%
```

由上面的数据可以得到以下结果

 - 每秒扫描约 25 个 url
 - 每秒约发送 2000 个请求
 - cpu 占用率约 60%

几乎是以前五倍的速度。

`max_paralle` 的解释见[插件配置](configration/plugins)页面，从上面的结果看到，即使到了 2000 qps，也还没有达到配置的检测最大并发，所以 `max_parallel` 一般情况下不需要修改。

`max_qps` 在非测试场景下，也很难打满，一般情况下只会改低，防止影响业务，而不需要再调高。

## 总结

限制 xray 扫描速度的主要原因为

 - xray 与被扫描目标之间的网络情况
 - 被扫描目标的性能；否则网络再好，扫描速度也无法提升

这两点可以从 `latency` 和  `failedRatio` 字段评估。`latency` 比较大，超过 100ms，说明网络比较慢，或者被扫描的目标响应比较慢。`failedRatio` 比较大，说明网络比较差，或者被扫描目标负载高而失去了响应等。

在网络情况稳定的情况下，可以考虑

 - 降低 `max_qps`，防止业务方负载过高
 - 关闭特定模块和禁用部分 poc，减少请求量
 - 不要性能比较差的代理作为 xray 的扫描代理，比如 burp 的代理通过请求的速度基本只有几十 qps，这样整个扫描的瓶颈就在 burp 上了。