# 检查更新

xray 内置了一个简单的更新检查机制，会在每次启动的时候检查有无新的版本发布，如果有更新将在界面上显示最新的 release notes。
如不需要该机制，可以通过下列方法禁用:

在 `config.yaml` 中添加如下配置即可禁用更新检查:

```yaml
update:
  check: false
```
