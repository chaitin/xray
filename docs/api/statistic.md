
## 字段说明:

- num_found_urls 发现的 url 数
- num_scanned_urls 扫描完成的 url 数
- num_sent_http_requests 已发送的 http 请求数
- average_response_time 最近 30s 平均响应时间
- ratio_failed_http_requests 最近 30s 请求失败率

控制台显示的 `pending` 就是 `num_found_urls - num_scanned_urls`，代表还没扫描的数量

## 样例
```json
{
  "num_found_urls": 0,
  "num_scanned_urls": 10,
  "num_sent_http_requests": 26,
  "average_response_time": 490.44446,
  "ratio_failed_http_requests": 0.26923078,
}
```


