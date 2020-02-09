# 使用 xray 进行服务扫描

xray 中最常见的是 web 扫描，但是 xray 将会逐渐开放服务扫描的相关能力，目前主要是服务扫描相关的 poc。老版本升级的用户请注意配置文件需要加入服务扫描的相关 poc 名字。

参数配置目前为 `target` 和 `json-output`。`target` 为 `host:port` 格式，比如 `127.0.0.1:8009`，`json-output` 同 web 扫描，将结果保存在 json 文件中。

```
NAME:
    servicescan - Run a service scan task

USAGE:
    servicescan [command options] [arguments...]

OPTIONS:
   --target value      specify the target, for example: host:8009
   --json-output FILE  output xray results to FILE in json format
```