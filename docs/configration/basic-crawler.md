基础爬虫的配置项对应于 `basic-crawler` 部分，默认的配置如下，用法参照文件中的注释:

```yaml
basic_crawler:
  max_depth: 0 # 爬虫最大深度, 0 为无限制
  max_count_of_links: 0 # 本次扫描总共爬取的最大连接数， 0 为无限制
  allow_visit_parent_path: false # 是否允许访问父目录, 如果扫描目标为 example.com/a/，若该项为 false, 那么就不会爬取 example.com/ 这级目录的内容
  restriction: # 和 mitm 中的写法一致, 有个点需要注意的是如果当前目标为 example.com 那么会自动添加 example.com 到 includes 中。
    includes: []
    excludes:
    - '*google*'
```
