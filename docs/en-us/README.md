<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> A powerful security assessment tool  🏠[Homepage](https://chaitin.github.io/xray/#/)  ⬇️[Download](https://github.com/chaitin/xray/releases) 📚[Chinese document](https://github.com/chaitin/xray)

### ✨ Demo

![](https://chaitin.github.io/xray/assets/term.svg)

## 🚀 Quick usage

1. Scan a single url
    
    ```bash
    xray webscan --url "http://example.com/?a=b"
    ```

1. Run as a HTTP proxy to scan passively
    
    ```bash
    xray webscan --listen 127.0.0.1:7777
    ```
    
   Configure the browser to use http proxy `http://127.0.0.1:7777`, then the proxy traffic can be automatically analyzed and scanned。

   >If need to scan https traffic，please read `capture https trafic` section in this document.

1. Specify the plugins to run manually
   
  By default, all built-in plugins are enabled, and the following commands can be used to enable specific plugins for this scan.
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. Specify plugin output path

    You can specify the output path of the vulnerability information:
    
    ```bash
    xray webscan --url http://example.com/?a=b --text-output result.txt
    xray webscan --url http://example.com/?a=b --json-output result.json
    ```

## 🛠 Detection module

We are working hard for new detection modules, for example xss detection, custom plugins, etc.

+ SQL injection detection (sqldet)
  
  Support for error based detection, boolean based detection, time based detection, and support mainstream databases.

+ Command injection detection (cmd_injection)

  Support for general command injection (shell), PHP code execution, template injection detection, and more.

+ Directory enumeration module (dirscan)

  Supports detection of more than 10 types of sensitive paths such as backup file leaks, temporary file leaks, debug pages, and configuration file leaks, covering most common cases.

+ Baseline detection (baseline)
  
  Detection of the ssl version of the remote host, http header, etc.

+ Arbitrary redirect (redirect)

  Support for redirection in html meta tag, 30x status code, etc.

+ Path traversal (path_traversal)

  Support multi-platform directory traversal vulnerabilities detection with automatically bypass technologies.

+ SSRF (ssrf)

  Support common SSRF vulnerabilities detection with automatically bypass technologies. This feature needs reverse server [reverse platform](https://chaitin.github.io/xray/#/guide/reverse).

+ CRLF injection (crlf_injection)

  Supports CRLF injection detection in header, query, and body positions.

+ JSONP sensitive information leak (jsonp)

  Support JSONP sensitive information leak vulnerabilities with smart detection algorithm.

+ ...


## ⚡️ Advanced usage

For the below advanced usage, please refer to [http://chaitin.github.io/xray/](http://chaitin.github.io/xray/) .

 - modify config file
 - generate ssl ca certificate
 - capture https traffic
 - modify http client config
 - the usage of reverse server
 - ...


## 📝 Discussion

If you have any questions, please post an issue on GitHub, or you can join the discussion groups below.

1. GitHub issue: https://github.com/chaitin/xray/issues
1. QQ group: 717365081
1. WeChat group: Scan the QR code below to add friends with me, I will invite you to the group.   

<img src="https://chaitin.github.io/xray/assets/wechat.jpg" height="150px">


