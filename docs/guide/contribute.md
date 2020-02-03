# 贡献 POC

xray 社区版经过数个版本的更迭，基本覆盖了对常见漏洞的 fuzzing, 稳定性和扫描效果上都有了很大的提高，在此感谢每一位为社区版反馈 bug 和要求功能改进的同学。但是漏洞扫描器这种的安全工具，POC 的质量和数量是一个不可绕过的鸿沟，而个人乃至团队的力量在这种事情面前也是微不足道的，因此，我们希望借助社区的力量来帮助 xray 更快更好的成长。一方面我们会通过大家的反馈对其不断进行迭代优化, 另一方面是集结大家的力量壮大它的 POC 库。

从以往经验来看，POC 宁缺毋滥。一个有问题的 POC 可能会影响到 xray 整体的运行。因此所以我们对社区提交的 POC 有着较为严格的约束。贡献者在提交 POC 的时，需要提供相应的可以复现的测试环境，测试环境以 docker 镜像的方式提供。至于难以做成 docker 镜像的漏洞，可以不提供。

## 流程

1. 贡献者以 PR 的方式向 github xray 社区仓库内提交 POC， 提交位置: [https://github.com/chaitin/xray/tree/master/pocs](https://github.com/chaitin/xray/tree/master/pocs)
2. PR 中根据 Pull Request 的模板填写 POC 信息
3. 内部审核 PR，确定是否合并入仓库
4. 每次发布新版时，CI 拉取 Github 仓库，并将社区的 POC 打包进社区版共享给大家

## POC 贡献规范

在提交之前请搜索仓库的 pocs 文件夹以及 Github 的 Pull request, 确保该 POC 没有被提交，而且不包含在部分扫描插件中（目前指 thinkPHP 和 struts 插件）。

1. 认真阅读 [如何编写高质量poc](guide/high_quality_poc.md) 文档
1. 期望近三年内主流框架，CMS等出现的漏洞，部分小众 CMS 的 POC 可能不被收录
1. 请直接保持单文件不要创建子目录； dockerfile 等文件不要放在仓库中，直接放在 PR 中即可
1. poc 请以 `.yml` 结尾，而不是 `.yaml`
1. poc name 一定是 `poc-yaml-` 开头，后面应该是 `[框架名/服务名/产品名等]-[cve编号]` 或者 `[框架名/服务名/产品名等]-[通用漏洞名称]`。比如 `elasticsearch-cve-2014-3120` 或者 `django-debug-page-info-leak`。无特殊情况，应该都是小写。poc 的 name 应和 yml 的文件名相同，比如上述 poc 的文件名应为 `django-debug-page-info-leak.yml`。poc name只能包含小写字母、短横线，版本号里的点号等符号请省略
1. poc 贡献者需要在 detail 中增加 author 字段，格式为 `name(link)`，name 可以为昵称，link 为可选项，一般使用个人 GitHub 首页或者博客链接等。
1. poc 贡献者需要在 detail 中增加 links 字段，这个字段的值是一个由URL组成的列表，表示和本漏洞和POC相关的参考链接，且一个POC至少需要有一个参考链接。这个链接可以是漏洞分析文章或测试环境地址等
1. 以上问题都可以通过 `./xray poclint` 来检查，具体使用请参考帮助文档
1. 测试环境可以参考 [vulhub](https://github.com/vulhub/vulhub/) [vulnapps](https://github.com/Medicean/VulApps) 。请勿直接填写公网上未修复的站点的地址，如果有特殊情况，请私聊解决。不接受没有测试环境的 poc
1. 一个 pull request 尽量只提交一个 poc，否则可能审核和修改过程会互相影响
1. 对于 0day / 1 day 等未大面积公开细节的漏洞请勿提交，可以私聊群管理员
1. 提交后，可以加一下我的微信 [点击查看](guide/feedback.md#反馈渠道) ，方便拉大家进群以及发放福利等

在 Github 提交 Pull request 后，会有travis-ci自动进行POC的格式检查，对于没有通过的检查，可以通过点击详情查看具体的错误原因：

![](../assets/contribute/fail-detail.jpg)

待修复完成并重新提交达到下面的样子就会被人工审核了。

![](../assets/contribute/pr.jpg)


