<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://chaitin.github.io/xray/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

> A powerful security assessment tool  🏠[Homepage](https://xray.cool)  ⬇️[Download](https://github.com/chaitin/xray/releases) 📚[Chinese document](https://github.com/chaitin/xray)

### ✨ Demo

![](https://chaitin.github.io/xray/assets/term.svg)

## 🚀 Quick usage

1. Use basic crawler to scan a website

    ```bash
    xray webscan --basic-crawler http://example.com --html-output crawler.html
    ```

1. Run as a HTTP proxy to scan passively
    
    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
    
   Configure the browser to use http proxy `http://127.0.0.1:7777`, then the proxy traffic can be automatically analyzed and scanned。

   >If need to scan https traffic，please read `capture https trafic` section in this document.
   
1. Scan a single url
    
    ```bash
    xray webscan --url http://example.com/?a=b --html-output single-url.html
    ```

1. Specify the plugins to run manually
   
  By default, all built-in plugins are enabled, and the following commands can be used to enable specific plugins for this scan.
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```
      
1. Specify plugin output path

    You can specify the output path of the vulnerability information:
    
    ```bash
    xray webscan --url http://example.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```

## 🛠 Detection module

We are working hard for new detection modules

 - xss

   XSS Vulnerabilities Scan

 - sqldet

   Support error based, boolean based and time based sql injection detection

 - cmd_injection

   Detect common shell command injection, PHP code execution and template injection, etc

 - dirscan

   Support about ten kinds of sensitive path and file type, including backup file, temp file, debug page, config file, etc

 - path_traversal

   Support command platform and encoding

 - xxe

   Support echo based detection and can work with reverse server

 - phantasm

   Common poc inside, user can add your own poc and run it. Document: https://chaitin.github.io/xray/#/guide/poc

 - upload

   Support common backend languages

 - brute_force

   The community version can detect weak password in http basic auth and simple form, common username and password dict inside

 - jsonp

   Detect jsonp api with sensitive data which can be called across origins

 - ssrf

   Support common bypass tech and can work with reverse server

 - baseline

   Detect outdated SSL version, missing or incorrect http headers, etc

 - redirect

   Detect arbitrary redirection from HTML meta and 30x response, etc

 - crlf_injection

   Detect CRLF injection in HTTP header, support parameters from query and body, etc

 - struts (Advanced version only)

   Check the target website for Struts2 series vulnerabilites, such as s2-016/s2-032/s2-045

 - thinkphp (Advanced version only)

   A plugin for ThinkPHP vulnerabilities

 - ...


## ⚡️ Advanced usage

For the below advanced usage, please refer to [https://chaitin.github.io/xray/](https://chaitin.github.io/xray/) .

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


