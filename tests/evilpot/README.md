Evil Pot
===

邪恶的罐子

一个专门用于让扫描器产生误报的靶场

编写插件应该尽量避免能在这个靶场扫描出结果

## 默认监听端口

- 8887: evil server 让扫描器产生误报 困难模式
    - 普通模式的基础上对所有请求元素进行拆解计算sha1/md5/base64
    - /etc/passwd和win.ini的内容
- 8888: evil server 让扫描器产生误报 普通模式
    - 常见状态码
    - 常见报错信息
    - 常见页面
    - 常见登录框
    - 常见xml头
    - 1-1000的sha1/md5/base64
    - 回显完整请求
    - 尝试计算请求中的算式
    - 尝试进行`sleep`和`wait for`的执行
- 8889: echo server 回显所有读到的数据

可在此处下载体验：https://github.com/chaitin/xray/releases?q=EvilPot&expanded=true