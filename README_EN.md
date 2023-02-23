<h1 align="center">Welcome to xray 👋</h1>
<p>
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://docs.xray.cool/#/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

[**中文版本**](./README.md)

> A powerful security assessment tool

### ✨ Demo

![](https://docs.xray.cool/assets/term.svg)

🏠[Document](https://docs.xray.cool)  ⬇️[Download](https://github.com/chaitin/xray/releases) 

Note: xray is not open source, you can download the built binaries directly, and the poc in the repository, which is mainly contributed by the community, will be packaged automatically for each xray release.

## 🚀 Quick use

**Before using the tool, read and agree to the terms especially the disclaimer in the [License](https://github.com/chaitin/xray/blob/master/LICENSE.md). Otherwise, do not install or use the tool.**

1. Use the base basic web crawler scan to scrape and perform vulnerability scanning on the scraped links

    ```bash
    xray webscan --basic-crawler http://example.com --html-output vuln.html
    ```

2. Use an HTTP proxy for passive scanning

    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
   Set the browser http proxy to `http://127.0.0.1:7777` to analyze proxy traffic and scan automatically.

   >To scan https traffic, read the `Crawling https Traffic` section below

3. Only a single url is scanned, without crawlers

    ```bash
    xray webscan --url http://example.com/?a=b --html-output single-url.html
    ```

4. Manually configure the plugins to run this time

   By default, all built-in plugins are enabled. You can specify which plugins are enabled for this scan using the following command.

   ```bash
   xray webscan --plugins cmd-injection,sqldet --url http://example.com
   xray webscan --plugins cmd-injection,sqldet --listen 127.0.0.1:7777
   ```

5. Specifying plugin output

   You can choose targeted vulnerability and their details of this scan and export a report:

    ```bash
    xray webscan --url http://example.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```

   [Sample Report](https://docs.xray.cool/assets/report_example.html)

For other uses, read the documentation： https://docs.xray.cool


## 🪟 Detection module

New detection modules will be added continuously

| 名称                                            | Key              | 版本                | 说明                                                                                                                                                                           |
|-----------------------------------------------|------------------|-------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| XSS vulnerability detection                   | `xss`            | Community Edition | Semantic analysis is used to detect XSS vulnerabilities                                                                                                                      |
| SQL Injection Detection                       | `sqldet`         | Community Edition | It supports error injection, Boolean injection, time blind injection, etc                                                                                                    |
| Command/Code injection detection              | `cmd-injection`  | Community Edition | It supports shell command injection, PHP code execution, template injection, etc                                                                                             |
| dirscan                                       | `dirscan`        | Community Edition | Detects more than 10 sensitive paths and files, including backup files, temporary files, debug pages, and configuration files                                                |
| Path traversal detection                      | `path-traversal` | Community Edition | Support for common platforms and encodings                                                                                                                                   |
| XML Entity Injection Detection                | `xxe`            | Community Edition | Support for echo and reverse platform detection                                                                                                                              |
| POC management                                | `phantasm`       | Community Edition | Some commonly used POCs are built in by default. Users can build and run POCs based on their needs. Document: [POC](https://docs.xray.cool/#/guide/poc)                      |
| File upload detection                         | `upload`         | Community Edition | Support for common back-end languages                                                                                                                                        |
| Weak password detection                       | `brute-force`    | Community Edition | The Community Edition supports the detection of HTTP basic authentication and weak passwords for simple forms, with a built-in dictionary of common user names and passwords |
| jsonp detection                               | `jsonp`          | Community Edition | Detects jsonp interfaces that contain sensitive information that can be read across domains                                                                                  |
| ssrf detection                                | `ssrf`           | Community Edition | ssrf detection module supports common bypass technology and reverse platform detection                                                                                       |
| Baseline detection                            | `baseline`       | Community Edition | Detects low SSL versions, missing or incorrectly added http, and so on                                                                                                       |
| Redirection detection                         | `redirect`       | Community Edition | Support HTML meta jump, 30x jump, etc                                                                                                                                        |
| CRLF injection                                | `crlf-injection` | Community Edition | Detects HTTP header injection and supports parameters for query, body, etc                                                                                                   |
| XStream vulnerability detection               | `xstream`        | Community Edition | Detect XStream series vulnerabilities                                                                                                                                        |
| Struts2 series vulnerability detection        | `struts`         | Advanced Edition  | Detect whether the target website has Struts2 series vulnerabilities, including s2-016, s2-032, s2-045 and other common vulnerabilities                                      |
| Thinkphp series vulnerability detection       | `thinkphp`       | Advanced Edition  | Detect vulnerabilities in ThinkPHP websites                                                                                                                                  |
| Shiro deserialization vulnerability detection | `shiro`          | Advanced Edition  | Detect Shiro deserialization vulnerability                                                                                                                                   |
| Fastjson series detection                     | `fastjson`       | Advanced Edition  | Detect fastjson series vulnerabilities                                                                                                                                       |


## ⚡️ Advanced use

See https://docs.xray.cool/ for the following advanced uses.

- Modifying the configuration file
- Scraping https traffic
- Modifying the http packet sending configuration
- Employing reverse platforms
- ...

## 😘 Contribute POC

The progress of xray is inseparable from the support of all masters. Upholding the spirit of mutual assistance and joint construction, in order to let us make progress together, xray has also opened the channel of "PoC collection"! Here you will get:

### Great Rewards

- Contributing to PoC earns generous points rewards and provides a fulfilling sense of accomplishment.
- Choose from over 50 types of peripheral gifts in our abundant gift redemption area.
- Regularly redeem JD Cards online and get one step closer to financial freedom.
- Get the opportunity to enter the core community, receive special tasks, and earn high rewards.

### Perfect Tutorial

- Perfect PoC writing tutorials and guidance help you get started quickly and avoid detours.

### Learning and Communication

- Interact face-to-face with contributors and developers to comprehensively improve various abilities.
- Get a direct interview opportunity without a written test, and turn your desired job into a reality.

If you have successfully contributed PoC but have not entered the contributor group, please add customer service WeChat:

<img src="https://docs.xray.cool/assets/customer_service.png?cache=_none" height="200px">

Provide your CT stack platform registration id for verification, and you can join the group after the verification is passed!

See: https://docs.xray.cool/#/guide/contribute

##  🔧 Surrounding ecology

### POC authoring aids

The tool can assist in generating POC, and the online version supports * * poc duplicate check * *, and the local version supports direct contract verification

#### Online version

- [**Rule Lab**](https://poc.xray.cool)
- The online version supports duplicate checking of **poc**

#### Local version

- [**gamma-gui**]( https://github.com/zeoxisca/gamma-gui)

### Xray gui aids

This tool is only a simple command line wrapper, not a direct method call. In the planning of xray, there will be a truly complete GUI version of XrayPro tool in the future. Please look forward to it.

- [**super-xray**]( https://github.com/4ra1n/super-xray)

## 📝 Discussion area

Feedback of  false positives,false negatives, please don’t hesitate to contact us. Be sure to read https://docs.xray.cool/#/guide/feedback first

If you have a question, you can ask it on GitHub or in the discussion group below

1. GitHub issue: https://github.com/chaitin/xray/issues
2. Wechat official account: Scan the following QR code on wechat and subscribe us

<img src="https://docs.xray.cool/assets/wechat.jpg?cache=_none" height="200px">

3. Wechat group: Please add the WeChat official account and click "联系我们" -> "加群", and then scan the QR code to join the group

4. QQ group: 717365081
