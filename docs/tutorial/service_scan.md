# 使用 xray 进行服务扫描

xray 中最常见的是 web 扫描，但是 xray 将会逐渐开放服务扫描的相关能力，目前主要是服务扫描相关的 poc。老版本升级的用户请注意配置文件需要加入服务扫描的相关 poc 名字，目前只有一个 tomcat-cve-2020-1938 ajp 协议任意文件检测 poc。

参数配置目前比较简单，输入支持两种方式，例如:

```
快速检测单个目标
./xray servicescan --target 127.0.0.1:8009

批量检查的 1.file 中的目标, 一行一个目标，带端口
./xray servicescan --target-file 1.file 

```

其中 1.file 的格式为一个行一个 service，如

```
10.3.0.203:8009
127.0.0.1:8009
```


也可以将结果输出到报告或json文件中

```
将检测结果输出到 html 报告中
./xray servicescan --target 127.0.0.1:8009 --html-output service.html
./xray servicescan --target-file 1.file --html-output service.html

将检测结果输出到 json 文件中
./xray servicescan --target 127.0.0.1:8099 --json-output 1.json 
```

```
NAME:
    servicescan - Run a service scan task

USAGE:
    servicescan [command options] [arguments...]

OPTIONS:
   --target value       specify the target, for example: host:8009
   --target-file value  load targets from a local file, one target a line
   --json-output FILE   output xray results to FILE in json format
   --html-output FILE   output xray result to `FILE` in HTML format
```

以 tomcat CVE-2020-1938 ajp 协议任意文件读取为例，命令行如下图。

![](../assets/tutorial/tomcat_servicescan.svg)
