**请先认真阅读下列要求，如不符合会被直接关闭 PR**

- 确保当前 POC 与已有的 POC 没有重复，除了仓库 `pocs` 目录中的，还有内置的几个用 Go 写的 POC也不要重复：
  ```
  poc-go-php-cve-2019-11043-rce
  poc-go-seeyon-htmlofficeservlet-rce
  poc-go-tongda-lfi-upload-rce
  poc-go-tongda-arbitrary-auth
  poc-go-ecology-dbconfig-info-leak
  poc-go-tomcat-put
  poc-go-tomcat-cve-2020-1938
  ```
  
- 阅读规范和要求 
  - https://chaitin.github.io/xray/#/guide/contribute
  - https://chaitin.github.io/xray/#/guide/high_quality_poc

- 一个 pull request 只提交一个 poc

- 对于 0day / 1 day 等未大面积公开细节的漏洞请勿提交，可以私聊群管理员

 - 不接受没有测试环境的 poc，测试环境可以使用 [vulhub](https://github.com/vulhub/vulhub/) 或 [vulnapps](https://github.com/Medicean/VulApps)，也可以提供 fofa/zoomeye/shodan 等搜索关键字。 请勿直接填写公网上未修复的站点的地址，如果有特殊情况，请私聊群内管理解决。
 
- 如果你的 poc 被合并或者没有合并但是评论说需要发送奖励，请查看 https://chaitin.github.io/xray/#/guide/feedback 并添加最下面的微信，说明你的 poc 地址，方便发送奖励。

**我是分割线，在提交 poc 填写说明的时候，请务必阅读上方要求，然后删除本分割线和上方的内容，只保留下面自定义的部分即可，否则不予通过。**

----------

## 本 poc 是检测什么漏洞的

## 测试环境

## 备注
