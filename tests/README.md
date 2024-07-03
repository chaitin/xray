# 测试靶场

## Evil Pot

`evilpot`目录下是我们实现的一个用于提高插件质量的测试靶场。

`evilpot`集成了一些常见的容易导致扫描器误报的情况，编写插件的过程应该尽量避免能在这个靶场扫描出结果。

## 常用靶场

这里依靠社区力量收集了几个常用的靶站，可以通过 docker/docker-compose 一键启用。活动开始后，我们注意到在这个repo
中 https://github.com/c0ny1/vulstudy ，作者已经收集了 12 个靶站，这些靶站基本符合我们的要求，所以后续提交中与该 repo
中靶站有所重复的将不再收录。

已有的靶站列表:

- DVWA
- bWAPP
- sqli-labs
- mutillidae
- BodgeIt
- WackoPicko
- WebGoat
- Hackademic
- XSSed
- DSVW
- vulnerable-node
- MCIR
