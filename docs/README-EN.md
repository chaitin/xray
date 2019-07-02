<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> A powerful security assessment tool  🏠 [Homepage](https://chaitin.github.io/xray/#/)  ⬇️ [Download](https://github.com/chaitin/xray/releases)


### ✨ Demo

![](https://chaitin.github.io/xray/assets/term.svg)

## 🚀 Quick usage

1. Scan single url
    
    ```bash
    xray webscan --url "http://example.com/?a=b"
    ```

1. Utilize HTTP proxy to scan passively
    
    ```bash
    xray webscan --listen 127.0.0.1:7777
    ```
   Set the http proxy of browser as `http://127.0.0.1:7777`, and the proxy traffic can be analyzed automatically and scan。
   
   >If need to scan https trafic，please read the below part of `capture https trafic` in this document

1. Manually specify the plugin to run this time
   
  By default, all built-in plugins will be enabled, and the following commands can be used to specify which plugins are enabled for this scan.
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. Specify output plugin

    You can specify to output the vulnerability information of this scan to a file:
    
    ```bash
    xray webscan --url http://example.com/?a=b --text-output result.txt
    xray webscan --url http://example.com/?a=b --json-output result.json
    ```

## 🛠 Detection module

New detection modules will continue to be added, and xss, custom plugins and other modules are on the road.

+ SQL injection detection(sqldet)
  
  Support for error injection, Boolean blinds, time blinds, and support for mainstream databases.

+ Command injection detection (cmd_injection)

  Support for general command injection (shell), PHP code execution, template injection detection, and more.

+ Directory enumeration module (dirscan)

  Supports detection of more than 10 types of sensitive paths such as backup file leaks, temporary file leaks, debug pages, and configuration file leaks, covering most common cases.

+ Baseline detection (baseline)
  
  Detection of the ssl version of the remote host, http header, etc.

+ Arbitrary redirect (redirect)

  Support for redirect of html meta in header.

+ Path traversal (path_traversal)

  Includes multi-platform payload to support detection of directory traversal vulnerabilities bypassed by common encodings.

+ SSRF (ssrf)

  Support no parity, prefix check, suffix check, and other partial bypass cases。This feature needs to combine [reverse platform](https://chaitin.github.io/xray/#/guide/reverse) to use.

+ CRLF injection (crlf_injection)

  Supports CRLF injection detection for header, query, and body positions.

+ JSONP sensitive information leak (jsonp)

  Built-in sensitive information analysis algorithms that detect jsonp vulnerabilities that can be exploited.

+ ...


## ⚡️ Advanced usage

For the below advanced usage, please refer to [http://chaitin.github.io/xray/](http://chaitin.github.io/xray/) 。

 - modify config file
 - generate certificate
 - capture https traffic
 - mofigy the config of https packets sending
 - the usage of reverse platform
 - ...


## 📝 Discussion

1. If you have any questions, you can post an issue on GitHub, or in the discussion group below.
1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ group: 717365081
1. Wechat group: Scan the QR code below and add my personal WeChat, which will pull everyone to the `xray` official WeChat group.   

<img src="https://chaitin.github.io/xray/assets/wechat.jpg" height="150px">
