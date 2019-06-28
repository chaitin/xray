<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> 一款功能强大的安全评估工具  🏠 [主页](https://chaitin.github.io/xray/#/)  ⬇️ [下载](https://github.com/chaitin/xray/releases)


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
   xray webscan --plugins cmd_injection,sqldet --proxy 127.0.0.1:7777
   ```
      
1. 指定插件输出

    可以指定将本次扫描的漏洞信息输出到某个文件中:
    
    ```bash
    xray webscan --url http://example.com/?a=b --output result.txt
    ```

## ⚡️ 进阶使用

下列高级用法请查看 [http://chaitin.github.io/xray/](http://chaitin.github.io/xray/) 使用。

 - 修改配置文件
 - 生成证书
 - 抓取 https 流量
 - 修改 https 发包配置
 - 反连平台的使用
 - ...


## 📝 讨论区

1. 如有问题可以在 GitHub 提 issue, 也可在下方的讨论组里
1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ 群: 717365081
1. 微信群: 扫描以下二维码加我的个人微信，会把大家拉到 `xray` 官方微信群    

<img src="https://chaitin.github.io/xray/assets/wechat.jpg" height="150px">