# 使用 xray 基础爬虫模式进行漏洞扫描

爬虫模式是模拟人工去点击网页的链接，然后去分析扫描，和代理模式不同的是，爬虫不需要人工的介入，访问速度要快很多，但是也有一些缺点需要注意

 - xray 的基础爬虫不能处理 js 渲染的页面，如果需要此功能，请参考 [版本对比](/generic/compare)
 - 需要首先人工配置登录 cookie，必需的 http 头等，如果登录失败，也不容易发现问题

## 启动爬虫


<!-- tabs:start -->

#### ** Windows **

```
./xray_windows_amd64 webscan --basic-crawler http://testphp.vulnweb.com/ --html-output xray-crawler-testphp.html
```


#### ** MacOS **

```
./xray_darwin_amd64 webscan --basic-crawler http://testphp.vulnweb.com/ --html-output xray-crawler-testphp.html
```

#### ** Linux **

```
./xray_linux_amd64 webscan --basic-crawler http://testphp.vulnweb.com/ --html-output xray-crawler-testphp.html
```

<!-- tabs:end -->