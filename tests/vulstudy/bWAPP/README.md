# 0x01 安装

```
# 安装docker
apt-get install docker.io
# 安装docker-compose
pip install docker-compose
# 下载vulstudy项目 
git clone https://github.com/c0ny1/vulstudy.git
```

## 0x02 使用
使用主要分两种：单独运行一个漏洞平台，同时运行多个漏洞平台。

#### 1.单独运行一个漏洞平台

cd到要运行的漏洞平台下运行以下命令

```
cd vulstudy/DVWA
docker-compose up -d #启动容器
docker-compose stop #停止容器
```

#### 2.同时运行所有漏洞平台

在项目根目录下运行以下命令

```
cd vulstudy
docker-compose up -d #启动容器
docker-compose stop #停止容器
```
![主界面](doc/vulstudy.jpg)

## 0x3 FAQ
**1.第一次启动bWAPP容器访问其主页会报错如下：**

```
Connection failed: Unknown database 'bWAPP'
```

**解决：** 第一次创建应事先访问/install.php来创建数据库！

**2.第一次搭建DVWA，在苹果系统下的safari浏览器下无法初始化数据库，并提示如下：**

```
CSRF token is incorrect
```

**解决：** 使用苹果系统下的其他浏览器即可，比如Chrome。

