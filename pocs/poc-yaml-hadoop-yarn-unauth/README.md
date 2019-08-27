## 使用环境

https://github.com/vulhub/vulhub/tree/master/confluence/CVE-2019-3396

## 验证截图

![1566371692813](./poc-yaml-Confluence-CVE-2019-3396-lfi.png)


## 备注
xray测试抓包返回的是json格式，故此建议只匹配javax.ws.rs.WebApplicationException