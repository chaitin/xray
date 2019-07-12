<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> 一款功能强大的安全评估工具  🏠[主页](https://chaitin.github.io/xray/#/)  ⬇️[下载](https://github.com/chaitin/xray/releases) 📚[English Document](https://github.com/chaitin/xray/tree/master/docs/en-us)

### ✨ Demo

![](https://chaitin.github.io/xray/assets/term.svg)

## 🚀 快速使用

1. 扫描单个 url
    
    ```bash
    xray webscan --url "http://example.com/?a=b"
    ```

1. 使用 HTTP 代理进行被动扫描
    
    ```bash
    xray webscan --listen 127.0.0.1:7777
    ```
   设置浏览器 http 代理为 `http://127.0.0.1:7777`，就可以自动分析代理流量并扫描。
   
   >如需扫描 https 流量，请阅读下方文档 `抓取 https 流量` 部分

1. 手动指定本次运行的插件
   
   默认情况下，将会启用所有内置插件，可以使用下列命令指定本次扫描启用的插件。
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. 指定插件输出

    可以指定将本次扫描的漏洞信息输出到某个文件中:
    
    ```bash
    xray webscan --url http://example.com/?a=b --text-output result.txt
    xray webscan --url http://example.com/?a=b --json-output result.json
    ```

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
    自定义 POC 请查看文档。
  

## 🛠 检测模块

新的检测模块将不断添加，xss，自定义插件等模块也在路上啦。

+ SQL 注入检测 (sqldet)
  
  支持报错注入、布尔盲注、时间盲注，支持主流数据库。

+ 命令注入检测 (cmd_injection)

  支持通用命令注入（shell）、PHP 代码执行、模板注入检测等。

+ 目录枚举模块 (dirscan)

  支持备份文件泄露、临时文件泄露、debug 页面、配置文件泄露等10余类敏感路径的检测，覆盖大多数常见的 case。

+ 基线检查 (baseline)
  
  对远程主机的 ssl 版本，http header 等的检测。

+ 任意跳转 (redirect)

  支持 html meta 跳转、30x 跳转等等。

+ 路径穿越 (path_traversal)

  包含多平台 payload，支持常见编码绕过的目录穿越漏洞的检测。

+ SSRF (ssrf)

  支持无校验情况、前缀校验情况、后缀校验情况和其他存在部分绕过情况。该功能需配合 [反连平台](https://chaitin.github.io/xray/#/guide/reverse) 使用。

+ CRLF 注入 (crlf_injection)

  支持 header, query, body 位置的 CRLF 注入检测。

+ JSONP 敏感信息泄露 (jsonp)

  内置敏感信息分析算法，能够检测到可以被利用的 jsonp 漏洞。

+ 自定义 POC 检测框架 (phantasm)

  默认内置部分常用的 poc，用户可以根据需要自行构建 poc 并运行，详情 [参考文档](https://chaitin.github.io/xray/#/guide/poc)

+ ...


## ⚡️ 进阶使用

下列高级用法请查看 [http://chaitin.github.io/xray/](http://chaitin.github.io/xray/) 使用。

 - 修改配置文件
 - 生成证书
 - 抓取 https 流量
 - 修改 https 发包配置
 - 反连平台的使用
 - ...


## 📝 讨论区

提交误报漏报需求等等请务必先阅读 [https://chaitin.github.io/xray/#/guide/feedback](https://chaitin.github.io/xray/#/guide/feedback)

如有问题可以在 GitHub 提 issue, 也可在下方的讨论组里

1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ 群: 717365081
1. 微信群: 扫描以下二维码加我的个人微信，会把大家拉到 `xray` 官方微信群    

<img src="https://chaitin.github.io/xray/assets/wechat.jpg" height="150px">
