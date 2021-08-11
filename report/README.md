# xray 漏洞模板

## DEMO

[报告样例](https://chaitin.github.io/xray/assets/report_example.html)

## 准备工作

安装 Node 和 yarn，进入该目录运行
```
yarn install
```

## 开发
```
yarn serve
```

## 构建
```
yarn build
```

## 注意

+ 为了方便报告在内网查看，在 build 时已经将所有依赖打包到 `dist/index.html`，漏洞模板只需用这一个文件，dist 目录下的其他文件可以忽略
+ 为了缩减体积，用了很多减少 bundle 大小的 tricks。其中对 antd icons 的处理稍微麻烦些，如果后续用到的组件有任何图标，都需要在 `src/icons` 手动引入，否则会显示不出来
+ 如果想自行将 json 漏洞数据输出为 html 报告，参考 `public/index.html` 中的注释内容, 注意需要做一些 html 转义，以防 xss 问题
