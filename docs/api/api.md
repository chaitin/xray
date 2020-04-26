
本章主要介绍 xray 内部使用的 api 格式，目前主要用于 `webhook-output` 或 `json-output` 的输出中。所有 api 均为 json 格式，其中 json 中的 `type` 项表明了当前的数据类型，比如:

web 漏洞的格式为:
```json
{
    "type": "web_vuln",
    "vuln":"xxxx",
}
```

web 统计类信息格式为：
```json
{
    "type": "web_statistic",
    "xxx": "xxx"
}
```

type 所有的类型为:

+ web_vuln
+ web_statistic
+ host_vuln (暂未开放)
+ host_statistic (暂未开放)