<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> 一款功能强大的安全评估工具  🏠[使用指南](https://chaitin.github.io/xray/#/)  ⬇️[下载](https://github.com/chaitin/xray/releases) 📚[English Document](https://github.com/chaitin/xray/tree/master/docs/en-us/generic)

### ✨ Demo

![](https://chaitin.github.io/xray/assets/term.svg)

## 🚀 快速使用

**在使用之前，请务必阅读并同意 [License](https://github.com/chaitin/xray/blob/master/LICENSE.md) 文件中的条款，否则请勿安装使用本工具。**

1. 使用基础爬虫爬取并对爬虫爬取的链接进行漏洞扫描
    
    ```bash
    xray webscan --basic-crawler http://example.com --html-output vuln.html
    ```

1. 使用 HTTP 代理进行被动扫描
    
    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
   设置浏览器 http 代理为 `http://127.0.0.1:7777`，就可以自动分析代理流量并扫描。
   
   >如需扫描 https 流量，请阅读下方文档 `抓取 https 流量` 部分

1. 只扫描单个 url，不使用爬虫
    
    ```bash
    xray webscan --url http://example.com/?a=b --html-output single-url.html
    ```

1. 手动指定本次运行的插件
   
   默认情况下，将会启用所有内置插件，可以使用下列命令指定本次扫描启用的插件。
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. 指定插件输出

    可以指定将本次扫描的漏洞信息输出到某个文件中:
    
    ```bash
    xray webscan --url http://example.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```
    
    [报告样例](https://chaitin.github.io/xray/assets/report_example.html)

1. 只运行单个**内置** POC

    在 xray 中，所有的 poc 隶属于插件 `phantasm`, 所以使用 poc 时需要开启 `phantasm` 插件才可生效。`--poc` 参数指定本次运行的 poc，如不指定，将运行所有的内置 poc。

    ```bash
    xray webscan --plugins phantasm --poc poc-yaml-thinkphp5-controller-rce --url http://example.com/
    ```
1. 运行用户自定义 POC

    用户可以按需书写自己的 YAML 格式的 POC， 并通过指定 `--poc` 参数运行，比如运行在 `/home/test/1.yaml` 处的 POC。

    ```bash
    xray webscan --plugins phantasm --poc /home/test/1.yaml --url http://example.com/
    ```
    
    `--poc` 参数非常灵活，支持 Glob 匹配，支持从目录加载，可以从以下几个例子理解用法:
    
    加载 `/home/test/pocs/` 所有的 POC:
    ```bash
    xray webscan --plugins phantasm --poc "/home/test/pocs/*"
    ```
    
    加载 `/home/test/pocs/` 下包含 thinkphp 的 POC
    ```bash
    xray webscan --plugins phantasm --poc "/home/test/pocs/*thinkphp*"
    ```
    
    自定义 POC 请查看文档。
  
1. 转发漏洞信息到数据库、邮件、IM 通知等

   用户可以使用 `--webhook-output` 将漏洞信息进行转发，后端需要返回 status 200 才认为发送成功，否则将打印错误日志。


   ```bash
   xray webscan --url http://example.com/ --webhook-output http://host:port/path
   ```


## 🛠 检测模块

新的检测模块将不断添加

 - XSS漏洞检测 (key: xss)

   利用语义分析的方式检测XSS漏洞

 - SQL 注入检测 (key: sqldet)

   支持报错注入、布尔注入和时间盲注等

 - 命令/代码注入检测 (key: cmd_injection)

   支持 shell 命令注入、PHP 代码执行、模板注入等

 - 目录枚举 (key: dirscan)

   检测备份文件、临时文件、debug 页面、配置文件等10余类敏感路径和文件

 - 路径穿越检测 (key: path_traversal)

   支持常见平台和编码

 - XML 实体注入检测 (key: xxe)

   支持有回显和反连平台检测

 - poc 管理 (key: phantasm)

   默认内置部分常用的 poc，用户可以根据需要自行构建 poc 并运行。文档：https://chaitin.github.io/xray/#/guide/poc

 - 文件上传检测 (key: upload)

   支持常见的后端语言

 - 弱口令检测 (key: brute_force)

   社区版支持检测 HTTP 基础认证和简易表单弱口令，内置常见用户名和密码字典

 - jsonp 检测 (key: jsonp)

   检测包含敏感信息可以被跨域读取的 jsonp 接口

 - ssrf 检测 (key: ssrf)

   ssrf 检测模块，支持常见的绕过技术和反连平台检测

 - 基线检查 (key: baseline)

   检测低 SSL 版本、缺失的或错误添加的 http 头等

 - 任意跳转检测 (key: redirect)

   支持 HTML meta 跳转、30x 跳转等

 - CRLF 注入 (key: crlf_injection)

   检测 HTTP 头注入，支持 query、body 等位置的参数
 
 - Struts2 系列漏洞检测 (高级版，key: struts)

   检测目标网站是否存在Struts2系列漏洞，包括s2-016、s2-032、s2-045等常见漏洞

 - Thinkphp系列漏洞检测 (高级版，key: thinkphp)

   检测ThinkPHP开发的网站的相关漏洞
 
 - ..


## ⚡️ 进阶使用

下列高级用法请查看 [http://chaitin.github.io/xray/](http://chaitin.github.io/xray/) 使用。

 - 修改配置文件
 - 生成证书
 - 抓取 https 流量
 - 修改 https 发包配置
 - 反连平台的使用
 - ...

## 贡献 POC

参照: [https://chaitin.github.io/xray/#/guide/contribute](https://chaitin.github.io/xray/#/guide/contribute)

## 📝 讨论区

提交误报漏报需求等等请务必先阅读 [https://chaitin.github.io/xray/#/guide/feedback](https://chaitin.github.io/xray/#/guide/feedback)

如有问题可以在 GitHub 提 issue, 也可在下方的讨论组里

1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ 群: 717365081
1. 微信群: 扫描以下二维码加我的个人微信，会把大家拉到 `xray` 官方微信群    

<img src="https://chaitin.github.io/xray/assets/wechat.jpg?_=1" height="200px">


