<h1 align="center">Welcome to xray 👋</h1>
<p align="center">
  <img src="https://img.shields.io/github/release/chaitin/xray.svg" />
  <img src="https://img.shields.io/github/release-date/chaitin/xray.svg?color=blue&label=update" />
  <img src="https://img.shields.io/badge/go report-A+-brightgreen.svg" />
  <a href="https://docs.xray.cool/">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" target="_blank" />
  </a>
</p>

<h3 align="center">A Powerful Security Assessment Tool </h3>

<p align="center">
  <a href="https://docs.xray.cool">🏠 Documentation</a> •
  <a href="https://github.com/chaitin/xray/releases">⬇️ Download xray</a> •
  <a href="https://github.com/chaitin/xpoc">⬇️ Download xpoc</a> •
  <a href="https://github.com/chaitin/xapp">⬇️ Download xapp</a> •
  <a href="https://github.com/chaitin/xray-plugins">📖 Plugin Repository</a>
</p>

[**中文版**](./README.md)

> Note: xray is not open source, just download the built binary files directly. The repository mainly contains community-contributed POCs. Each xray release will automatically package them.

## ✨ xray 2.0

To address the issue of xray 1.0 becoming complex and bloated with added features, we have launched xray 2.0.

This new version aims to enhance the smoothness of functionality, lower the usage threshold, and help more security industry practitioners achieve a better experience more efficiently. xray 2.0 will integrate a series of new security tools to form a comprehensive security toolset.

**The second tool in the xray 2.0 series, xapp, is now online, welcome to try it!**

### XPOC

xpoc is the first tool in the xray 2.0 series, designed as a rapid emergency response tool for supply chain vulnerability scanning.

Project address: https://github.com/chaitin/xpoc

### XAPP

xapp is a tool focused on web fingerprint identification. You can use xapp to identify the technologies used by web targets and prepare for security testing.

Project address: https://github.com/chaitin/xapp

### Plugin Repository

We have created a dedicated repository for various plugins, aimed at facilitating the sharing and use of different plugins.

It mainly collects open-source scripts converted into xray format for everyone to use.

We will periodically push some new plugins here and hope everyone will actively optimize or submit plugins to enrich this repository together.

Project address: https://github.com/chaitin/xray-plugins

## 🚀 Quick Usage

**Before using, be sure to read and agree to the terms in the [License](https://github.com/chaitin/xray/blob/master/LICENSE.md) file. If not, please do not install or use this tool.**

1. Use the basic crawler to scan the links crawled by the crawler for vulnerabilities

    ```bash
    xray webscan --basic-crawler http://example.com --html-output vuln.html
    ```

2. Use HTTP proxy for passive scanning

    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
   Set the browser's HTTP proxy to `http://127.0.0.1:7777`, then you can automatically analyze proxy traffic and scan it.

   > To scan HTTPS traffic, please read the "Capture HTTPS Traffic" section below.

3. Scan a single URL without using a crawler

    ```bash
    xray webscan --url http://example.com/?a=b --html-output single-url.html
    ```

4. Manually specify plugins for this run

   By default, all built-in plugins will be enabled. You can specify the plugins to be enabled for this scan with the following commands.

   ```bash
   xray webscan --plugins cmd-injection,sqldet --url http://example.com
   xray webscan --plugins cmd-injection,sqldet --listen 127.0.0.1:7777
   ```

5. Specify Plugin Output

   You can specify to output the vulnerability information of this scan to a file:

    ```bash
    xray webscan --url http://example.com/?a=b \
    --text-output result.txt --json-output result.json --html-output report.html
    ```

   [Sample Report](https://docs.xray.cool/assets/report_example.html)

For other usage, please read the documentation: https://docs.xray.cool

## 🪟 Detection Modules

New detection modules will be continuously added.

| Name             | Key              | Version  | Description                                                                          |
|------------------|------------------|----------|--------------------------------------------------------------------------------------|
| XSS Detection    | `xss`            | Community | Detects XSS vulnerabilities using semantic analysis                                  |
| SQL Injection Detection | `sqldet` | Community | Supports error-based injection, boolean-based injection, and time-based blind injection |
| Command/Code Injection Detection | `cmd-injection`  | Community | Supports shell command injection, PHP code execution, template injection, etc.        |
| Directory Enumeration | `dirscan`    | Community | Detects over 10 types of sensitive paths and files such as backup files, temporary files, debug pages, and configuration files |
| Path Traversal Detection | `path-traversal` | Community | Supports common platforms and encodings                                               |
| XML External Entity (XXE) Detection | `xxe`    | Community | Supports detection with echo and back-connect platform                                 |
| POC Management      | `phantasm`       | Community | Comes with some common POCs by default; users can build and run POCs as needed. Documentation: [POC](https://docs.xray.cool/#/guide/poc) |
| File Upload Detection | `upload` | Community | Supports common backend languages                                                     |
| Weak Password Detection | `brute-force` | Community | Community edition supports HTTP basic authentication and simple form weak password detection, with built-in common username and password dictionary |
| JSONP Detection      | `jsonp` | Community | Detects JSONP interfaces containing sensitive information that can be read across domains |
| SSRF Detection      | `ssrf` | Community | SSRF detection module, supports common bypass techniques and back-connect platform detection |
| Baseline Check      | `baseline` | Community | Detects low SSL versions, missing or incorrectly added HTTP headers                   |
| Arbitrary Redirect Detection | `redirect` | Community | Supports HTML meta redirects, 30x redirects, etc.                                     |
| CRLF Injection      | `crlf-injection` | Community | Detects HTTP header injection, supports parameters in query, body, etc.               |
| XStream Vulnerability Detection | `xstream` | Community | Detects XStream series vulnerabilities                                                |
| Struts2 Vulnerability Detection | `struts` | Advanced | Detects Struts2 series vulnerabilities, including common ones like s2-016, s2-032, s2-045, s2-059, s2-061, etc. |
| ThinkPHP Vulnerability Detection | `thinkphp` | Advanced | Detects related vulnerabilities in websites developed with ThinkPHP                   |
| Shiro Deserialization Vulnerability Detection | `shiro` | Advanced | Detects Shiro deserialization vulnerabilities                                         |
| Fastjson Vulnerability Detection | `fastjson` | Advanced | Detects Fastjson series vulnerabilities                                               |


## ⚡️ Advanced Usage

For the following advanced usage, please see https://docs.xray.cool/.

- Modify configuration file
- Capture HTTPS traffic
- Modify HTTP request configuration
- Use of the back-connect platform
- ...

## 😘 Contribute POCs

The progress of xray cannot be achieved without the support of many contributors. In the spirit of mutual assistance and co-construction, to help us all progress together, xray has also opened a channel for "POC inclusion"! Here you will get:

### Submission Process

1. Contributors submit to the GitHub xray community repository via PR. POC submission location: https://github.com/chaitin/xray/tree/master/pocs, fingerprint script submission location: https://github.com/chaitin/xray/tree/master/fingerprints
2. Fill in the POC information according to the Pull Request template in the PR.
3. Internal review of the PR to determine whether to merge into the repository.
4. Note: To receive POC rewards, you need to submit your POC to the CT stack to receive the rewards.

### Generous Rewards

- Contribute POCs to receive **generous gold rewards** with a sense of accomplishment;
- **Rich gift** redemption area, with over 50 kinds of peripheral gifts to choose from;
- Regularly launch JD card redemption, getting **closer to financial freedom**;
- Opportunity to enter the core community, receive special tasks, and earn **high bounties**;

### Comprehensive Tutorials

- Comprehensive **POC writing tutorials and guidance** to help you get started quickly and avoid pitfalls;

### Learning and Communication

- **Face-to-face learning and communication opportunities** with contributors and developers, improving comprehensive skills;
- **Direct interview opportunities** without written tests, making good jobs not just a dream;

If you have successfully contributed a POC but have not yet

joined the group, please add the customer service WeChat:

<img src="./asset/customer_service.png" height="200px">

Provide the platform registration ID for verification. Once verified, you can join the group!

Refer to: https://docs.xray.cool/#/guide/contribute

## 🔧 Surrounding Ecosystem


### POC Quality Confirmation Range

[**Evil Pot**](https://github.com/chaitin/xray/tree/master/tests/evilpot)

[Releases](https://github.com/chaitin/xray/releases?q=EvilPot&expanded=true)

A range specifically designed to allow scanners to generate false positives

Plugins should be written to try to avoid being able to scan results in this range

### POC Writing Assistant Tools

This tool can assist in generating POCs, and the online version supports **POC duplication checks**, while the local version supports direct packet verification.

#### Online Version
- [**Rule Laboratory**](https://poc.xray.cool)
- The online version supports **POC duplication checks**.
#### Local Version
- [**gamma-gui**](https://github.com/zeoxisca/gamma-gui)

### xray GUI Assistant Tools

This tool is just a simple command line wrapper, not a direct method call. In the xray plan, there will be a truly complete GUI version of XrayPro in the future, so stay tuned.

- [**super-xray**](https://github.com/4ra1n/super-xray)

## 📝 Discussion Area

Fellow developers and xray fans, feel free to come to [Discussion Board Vote](https://github.com/chaitin/xray/discussions/1804) to decide the development priorities for xray 2.0 tools and let your voice shape the future of xray! 🚀

Before submitting false positives, missed reports, or requests, please be sure to read https://docs.xray.cool/#/guide/feedback.

If you have any questions, you can submit an issue on GitHub or join the discussion groups below.

1. GitHub:
   - https://github.com/chaitin/xray/issues
   - https://github.com/chaitin/xray/discussions

2. WeChat Official Account: Scan the QR code below to follow us

    <img src="./asset/wechat.jpg" height="200px">

3. WeChat Group: Add the WeChat official account, click "Contact Us" -> "Join Group", and then scan the QR code to join the group.

4. QQ Group: 717365081

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=chaitin/xray&type=Date)](https://star-history.com/#chaitin/xray&Date)