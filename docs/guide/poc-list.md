# 内置POC列表

社区版XRay将内置一些常用漏洞的POC，在XRay每次更新后，可以删除本地配置文件，重新生成最新的配置文件。此时，你将看到所有内置POC：

```yaml
  phantasm:
    enabled: true
    max_parallel: 1
    poc:
    - poc-yaml-drupal-drupalgeddon2-rce
    - poc-yaml-joomla-cve-2015-7297-sqli
    - poc-yaml-joomla-cve-2017-8917-sqli
    - poc-yaml-thinkphp5-controller-rce
    - poc-yaml-thinkphp5023-method-rce
    - poc-go-tomcat-put
    - ...
```

`phantasm.poc`是一个列表，里面包含所有内置POC，列举如下：

- `poc-yaml-drupal-drupalgeddon2-rce` Drupal远程代码执行漏洞（CVE-2018-7600）
- `poc-yaml-joomla-cve-2015-7297-sqli` Joomla SQL注入漏洞（CVE-2015-7297）
- `poc-yaml-joomla-cve-2017-8917-sqli` Joomla SQL注入漏洞（CVE-2017-8917）
- `poc-yaml-thinkphp5-controller-rce` ThinkPHP < 5.0.23 远程代码执行漏洞
- `poc-yaml-thinkphp5023-method-rce` ThinkPHP 5.0/5.1 远程代码执行漏洞
- `poc-go-tomcat-put` Tomcat PUT 文件写入漏洞（CVE-2017-12615）
