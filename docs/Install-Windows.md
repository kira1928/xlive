# Install Windows

## Step 1: 获取ffmpeg
从[FFmpeg Builds](https://www.gyan.dev/ffmpeg/builds/#release-builds)下载压缩包，将其中`bin/ffmpeg.exe`解压出来备用

## Step 2: 下载 xlive
打开[xlive Releases](https://github.com/kira1928/xlive/releases/latest)
- 32位系统下载`xlive-windows-386.zip`，并解压
- 64位系统下载`xlive-windows-amd64.zip`，并解压  

之后将`ffmpeg.exe`复制到和`xlive`同一目录下

## Step 3: 运行
双击 xlive-windows-amd64.exe 执行将使用和 exe 文件在同一目录下的 `config.yml` 文件作为默认配置文件启动程序。

也可以执行
```
./xlive-windows-amd64 -c ./config.yml
```
使用指定的配置文件来启动