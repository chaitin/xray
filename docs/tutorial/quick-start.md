# 快速使用

这里提供一些简单的使用方法参考，高阶用户可以直接 `xray -h` 通过查看命令行参数使用。

> 在介绍页中讲到了 [xray 的简易架构设计](basic/introduce.md#简易架构)，推荐先阅读一下该节内容，以便更好的使用 xray。

以下的 xray 指代下载的二进制程序，如果是 windows 平台请更换为 `xray_windows_amd64.exe` 或 `xray_windows_386.exe`，其他平台以此类推。

<!-- tabs:start -->

#### ** English **

Hello!

#### ** French **

Bonjour!

#### ** Italian **

Ciao!

<!-- tabs:end -->

1. 使用基础爬虫爬取并对爬虫爬取的链接进行漏洞扫描
    
    ```bash
    xray webscan --basic-crawler http://example.com --html-output vuln.html
    ```

1. 使用 HTTP 代理进行被动扫描
    
    ```bash
    xray webscan --listen 127.0.0.1:7777 --html-output proxy.html
    ```
   设置浏览器 http 代理为 `http://127.0.0.1:7777`，然后使用浏览器访问网页，就可以自动分析代理流量并扫描。
   
   > 如需扫描 https 流量，请阅读文档 `抓取 https 流量` 部分

1. 快速测试单个 url, **无爬虫**
    
    ```bash
    xray webscan --url http://example.com/?a=b --html-output single-url.html
    ```

1. 手动指定本次运行的插件
   
   默认情况下，将会启用所有内置插件，可以使用下列命令指定本次扫描启用的插件。
   
   ```bash
   xray webscan --plugins cmd_injection,sqldet --url http://example.com
   xray webscan --plugins cmd_injection,sqldet --listen 127.0.0.1:7777
   ```

在这里介绍了 xray 的常用的几个命令，至于更高级的用法和参数请继续阅读下一章 [配置文件](../configration/cli.md)