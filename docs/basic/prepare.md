# 下载运行

xray 为单文件二进制文件，无依赖，也无需安装，下载后直接使用。

## 下载地址

请下载的时候选择最新的版本下载。

+ Github: https://github.com/chaitin/xray/releases
+ 网盘: https://yunpan.360.cn/surl_y3Gu6cugi8u

> 注意： 不要直接 clone 仓库，xray 并不开源，仓库内不含源代码，直接下载构建的二进制文件即可。

## 选择版本

xray 跨平台支持，请下载时选择需要的版本下载。

![image](https://user-images.githubusercontent.com/20637881/66907963-dc863e00-f03c-11e9-8de9-6a79835dbd15.png)

对于上述图片中的文件，说明如下:

+ `darwin_amd64` MacOS
+ `linux_386` Linux x86
+ `linux_amd64` Linux x64
+ `windows_386` Windows x86
+ `windows_amd64` Windows x64 
+ `sha256.txt`

    校验文件，内含个版本的 sha256 的哈希值，请下载后自行校验以防被劫持投毒。
+ `Source Code` 为 Github 自动打包的，无意义，请忽略。

## 运行

下载对应系统的版本后，解压 zip 文件，Linux/Mac 用户在终端 (Terminal) 运行， Windows 用户请在 Powershell 或其他高级 Shell 中运行，在 CMD 中运行可能体验不佳。

如果一切顺利，你可以通过运行 `version` 命令查看版本:

```
$ ./xray_darwin_amd64 version

 __   __  _____              __     __
 \ \ / / |  __ \      /\     \ \   / /
  \ V /  | |__) |    /  \     \ \_/ /
   > <   |  _  /    / /\ \     \   /
  / . \  | | \ \   / ____ \     | |
 /_/ \_\ |_|  \_\ /_/    \_\    |_|


Version: 0.14.0/d1742479/COMMUNITY

Generate default configurations to config.yaml
[xray 0.14.0/d1742479]
Build: [2019-10-16] [darwin/amd64] [RELEASE/COMMUNITY]
Compiler Version: go version go1.13.1 linux/amd64

To show open source licenses, please use `osslicense` sub-command.
```

接下来就可以体验社区最强大的安全监测工具了。