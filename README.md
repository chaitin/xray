<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://docs.xray.cool/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

[**English Version**](./README_EN.md)

> 一款功能强大的安全评估工具 

## ✨ Demo

![](https://docs.xray.cool/assets/term.svg)

🏠[使用文档](https://docs.xray.cool)  
⬇️[国内用户下载地址](https://stack.chaitin.com/tool/detail?id=1)  
⬇️[GitHub下载地址](https://github.com/chaitin/xray/releases)

> 注意：xray 不开源，直接下载构建的二进制文件即可，仓库内主要为社区贡献的 poc，每次 xray 发布将自动打包。

## xray2.0

为了解决 xray 1.0在功能增加过程中变得复杂且臃肿的问题，我们推出了 xray 2.0。

这一全新版本致力于提升功能使用的流畅度，降低使用门槛，并帮助更多安全行业从业者以更高效的模式收获更好的体验。xray 2.0 将整合一系列新的安全工具，形成一个全面的安全工具集。

**xray2.0系列的第一款工具xpoc已经上线，欢迎体验！**

- [**xpoc**](https://github.com/chaitin/xpoc)

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
   xray webscan --plugins cmd-injection,sqldet --url http://example.com
   xray webscan --plugins cmd-injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. 指定插件输出

    可以指定将本次扫描的漏洞信息输出到某个文件中:
    
    ```bash
    xray webscan --url http://example.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```
    
    [报告样例](https://docs.xray.cool/assets/report_example.html)

其他用法请阅读文档： https://docs.xray.cool

## 🪟 检测模块

新的检测模块将不断添加

| 名称             | Key              | 版本  | 说明                                                                              |
|----------------|------------------|-----|---------------------------------------------------------------------------------|
| XSS漏洞检测        | `xss`            | 社区版 | 利用语义分析的方式检测XSS漏洞                                                                |
| SQL 注入检测       | `sqldet`         | 社区版 | 支持报错注入、布尔注入和时间盲注等                                                               |
| 命令/代码注入检测      | `cmd-injection`  | 社区版 | 支持 shell 命令注入、PHP 代码执行、模板注入等                                                    |
| 目录枚举           | `dirscan`        | 社区版 | 检测备份文件、临时文件、debug 页面、配置文件等10余类敏感路径和文件                                           |
| 路径穿越检测         | `path-traversal` | 社区版 | 支持常见平台和编码                                                                       |
| XML 实体注入检测     | `xxe`            | 社区版 | 支持有回显和反连平台检测                                                                    |
| poc 管理         | `phantasm`       | 社区版 | 默认内置部分常用的 poc，用户可以根据需要自行构建 poc 并运行。文档：[POC](https://docs.xray.cool/#/guide/poc) |
| 文件上传检测         | `upload`         | 社区版 | 支持常见的后端语言                                                                       |
| 弱口令检测          | `brute-force`    | 社区版 | 社区版支持检测 HTTP 基础认证和简易表单弱口令，内置常见用户名和密码字典                                          |
| jsonp 检测       | `jsonp`          | 社区版 | 检测包含敏感信息可以被跨域读取的 jsonp 接口                                                       |
| ssrf 检测        | `ssrf`           | 社区版 | ssrf 检测模块，支持常见的绕过技术和反连平台检测                                                      |
| 基线检查           | `baseline`       | 社区版 | 检测低 SSL 版本、缺失的或错误添加的 http 头等                                                    |
| 任意跳转检测         | `redirect`       | 社区版 | 支持 HTML meta 跳转、30x 跳转等                                                         |
| CRLF 注入        | `crlf-injection` | 社区版 | 检测 HTTP 头注入，支持 query、body 等位置的参数                                                |
| XStream漏洞检测    | `xstream`        | 社区版 | 检测XStream系列漏洞                                                                   |
| Struts2 系列漏洞检测 | `struts`         | 高级版 | 检测目标网站是否存在Struts2系列漏洞，包括s2-016、s2-032、s2-045、s2-059、s2-061等常见漏洞                 |
| Thinkphp系列漏洞检测 | `thinkphp`       | 高级版 | 检测ThinkPHP开发的网站的相关漏洞                                                            |
| shiro反序列化漏洞检测  | `shiro`          | 高级版 | 检测Shiro反序列化漏洞                                                                   |
| fastjson系列检测   | `fastjson`       | 高级版 | 检测fastjson系列漏洞                                                                  |


## ⚡️ 进阶使用

下列高级用法请查看 https://docs.xray.cool/ 使用。

 - 修改配置文件
 - 抓取 https 流量
 - 修改 http 发包配置
 - 反连平台的使用
 - ...

## 😘 贡献 POC

xray的进步离不开各位师傅的支持，秉持着互助共建的精神，为了让我们共同进步，xray也开通了“PoC收录”的渠道！在这里你将会得到：

### 提交流程

1. 贡献者以 PR 的方式向 github xray 社区仓库内提交， POC 提交位置: https://github.com/chaitin/xray/tree/master/pocs, 指纹识别脚本提交位置: https://github.com/chaitin/xray/tree/master/fingerprints
2. PR 中根据 Pull Request 的模板填写 POC 信息
3. 内部审核 PR，确定是否合并入仓库
4. 但需要注意，如果想要获得POC的奖励，需要将你的POC提交到CT stack，才能获取到奖励

### 丰厚的奖励

- 贡献PoC将获得**丰厚的金币奖励**，成就感满满；
- **丰富的礼品**兑换专区，50余种周边礼品任你挑选；
- 定期更有京东卡上线兑换，离**财富自由**又近了一步；
- 进入核心社群的机会，领取特殊任务，赚取**高额赏金**；

### 完善的教程

- 完善的**PoC编写教程和指导**，让你快速上手，少走弯路；

### 学习与交流

- **与贡献者、开发者面对面**学习交流的机会，各项能力综合提高；
- 免笔试的**直通面试机会**，好工作不是梦；

如果你已经成功贡献过PoC但是还没有进群，请添加客服微信：

<img src="https://docs.xray.cool/assets/customer_service.png?cache=_none" height="200px">

提供平台注册id进行验证，验证通过后即可进群！

参照: https://docs.xray.cool/#/guide/contribute

## 🔧周边生态

### POC编写辅助工具

该工具可以辅助生成POC，且在线版支持**poc查重**，本地版支持直接发包验证

#### 在线版
- [**规则实验室**](https://poc.xray.cool)
- 在线版支持对**poc查重**
#### 本地版
- [**gamma-gui**](https://github.com/zeoxisca/gamma-gui)

### xray gui辅助工具

本工具仅是简单的命令行包装，并不是直接调用方法。在 xray 的规划中，未来会有一款真正的完善的 GUI 版 XrayPro 工具，敬请期待。

- [**super-xray**](https://github.com/4ra1n/super-xray)

## 📝 讨论区

提交误报漏报需求等等请务必先阅读 https://docs.xray.cool/#/guide/feedback

如有问题可以在 GitHub 提 issue, 也可在下方的讨论组里

1. GitHub issue: https://github.com/chaitin/xray/issues

2. 微信公众号：微信扫描以下二维码，关注我们

<img src="https://docs.xray.cool/assets/wechat.jpg?cache=_none" height="200px">

3. 微信群: 请添加微信公众号并点击“联系我们" -> "加群“，然后扫描二维码加群

4. QQ 群: 717365081


