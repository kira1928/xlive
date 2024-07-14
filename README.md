# xlive

xlive 是一个支持多种直播平台的直播录制工具。本项目 fork 自 [Bililive-go](https://github.com/hr3lxphr6j/bililive-go)。   

![image](docs/screenshot.webp)

## 支持网站

<table>
    <tr align="center">
        <th>站点</th>
        <th>url</th>
        <th>cookie</th>
    </tr>
    <tr align="center">
        <td>Acfun直播</td>
        <td>live.acfun.cn</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>哔哩哔哩直播</td>
        <td>live.bilibili.com</td>
        <td>支持</td>
    </tr>
    <tr align="center">
        <td>战旗直播</td>
        <td>www.zhanqi.tv</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>斗鱼直播</td>
        <td>www.douyu.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>火猫直播</td>
        <td>www.huomao.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>龙珠直播</td>
        <td>longzhu.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>虎牙直播</td>
        <td>www.huya.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>CC直播</td>
        <td>cc.163.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>一直播</td>
        <td>www.yizhibo.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>OPENREC</td>
        <td>www.openrec.tv</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>企鹅电竞</td>
        <td>egame.qq.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>浪live</td>
        <td>play.lang.live & www.lang.live</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>花椒</td>
        <td>www.huajiao.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>抖音直播</td>
        <td>live.douyin.com</td>
        <td>支持</td>
    </tr>
    <tr align="center">
        <td>猫耳</td>
        <td>fm.missevan.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>克拉克拉</td>
        <td>www.hongdoufm.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>YY直播</td>
        <td>www.yy.com</td>
        <td></td>
    </tr>
    <tr align="center">
        <td>微博直播</td>
        <td>weibo.com</td>
        <td></td>
    </tr>
</table>

### cookie 在 config.yml 中的设置方法

cookie的设置以域名为单位。比如想在录制抖音直播时使用 cookie，那么 `config.yml` 中可以像下面这样写：
```
cookies:
  live.douyin.com: __ac_nonce=123456789012345678903;name=value
```
这里 name 和 value 只是随便举的例子，用来说明当添加超过一条 cookie 的键值对时应该用分号隔开。
至于具体应该添加哪些键，就需要用户针对不同网站自己获取了。

## 在网页中修改设置

点击网页左边的 `设置` 可以在线修改项目的配置文件，之后点击页面下面的 `保存设置` 按钮保存设置。
如果保存后窗口提醒设置保存成功，那就是配置文件已经被写入磁盘了。如果是保存失败，那可能是配置文件格式问题或者遇到程序 bug，总之磁盘上的配置文件没变。

在网页中即使保存配置成功也不一定表示相应的配置会立即生效。
有些配置需要停止监控后再重新开始监控才会生效，有些配置也许要重启程序才会生效。

## 网页播放器

点击对应直播间行右边的 `文件` 链接可以跳转到对应直播间的录播目录中。  
当然你点左边的 `文件` 一路找过去也行。

https://github.com/hr3lxphr6j/bililive-go/assets/2352900/6453900c-6321-417b-94f2-d65ec2ab3d7e

## 依赖
* [ffmpeg](https://ffmpeg.org/)

## 安装和使用

[Windows](docs/Install-Windows.md)

[macOS](docs/Install-macOS.md)

[Linux](docs/Install-Linux.md)

### docker

使用 https://hub.docker.com/r/kira1928/xlive 镜像创建容器运行。

例如：
```
docker run --restart=always -v ./config.yml:/etc/xlive/config.yml -v ./Videos:/srv/xlive -p 8080:8080 -d kira1928/xlive
```

### docker compose

使用项目根目录下的 `docker-compose.yml` 配置文件启动 docker compose 运行。

例如：
```
docker compose up
```
此时默认使用 `config.docker.yml` 文件作为程序的配置文件，`Videos/` 目录作为录制视频的输出目录。

NAS 用户使用系统自带 GUI 创建 docker compose 的情况请参考群晖用 docker compose 安装 bgo 的 [图文说明](./docs/Synology-related.md#如何用-docker-compose-安装-bgo)

## 常见问题
[docs/FAQ.md](docs/FAQ.md)

## 开发环境搭建

### 一、环境准备

#### 1. 前端环境

1）前往 https://nodejs.org/zh-cn/ 下载当前版本node（v20.15.1）

2）命令行运行 `node -v` 若控制台输出版本号则前端环境搭建成功

#### 2.后端环境
1）下载golang安装 版本号1.22

- 国际: https://golang.org/dl/
- 国内: https://golang.google.cn/dl/

2）命令行运行 go 若控制台输出各类提示命令 则安装成功 输入 `go version` 确认版本

### 二、克隆代码并编译
克隆：
```
git clone https://github.com/kira1928/xlive.git
cd xlive
```

使用 make 命令编译：
```
make build-web
make
```
没有安装 make 的话，可以直接执行以下命令代替：
```
go run build.go build-web
go run build.go release
```


### 三、编译 debug 版本
使用 make 命令编译：
```
make dev
```
没有安装 make 的话，可以直接执行以下命令代替：
```
go run build.go dev
```

## 其他文档
- [输出文件名模板](docs/File-name-template.md)
- [API doc](docs/API.md)
- [Grafana 面板](docs/grafana.md)

## 参考
- [you-get](https://github.com/soimort/you-get)
- [ykdl](https://github.com/zhangn1985/ykdl)
- [youtube-dl](https://github.com/ytdl-org/youtube-dl)
