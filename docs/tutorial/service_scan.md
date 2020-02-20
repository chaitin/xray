# 使用 xray 进行服务扫描

xray 中最常见的是 web 扫描，但是 xray 将会逐渐开放服务扫描的相关能力，目前主要是服务扫描相关的 poc。老版本升级的用户请注意配置文件需要加入服务扫描的相关 poc 名字，目前只有一个 tomcat-cve-2020-1938 ajp 协议任意文件检测 poc。

参数配置目前比较简单，为 `target` 和 `json-output`。`target` 为 `host:port` 格式，比如 `127.0.0.1:8009`，`json-output` 同 web 扫描，将结果保存在 json 文件中。

```
NAME:
    servicescan - Run a service scan task

USAGE:
    servicescan [command options] [arguments...]

OPTIONS:
   --target value      specify the target, for example: host:8009
   --json-output FILE  output xray results to FILE in json format
```

以 tomcat CVE-2020-1938 ajp 协议任意文件读取为例，命令行如下图。

![](../assets/tutorial/tomcat_servicescan.svg)