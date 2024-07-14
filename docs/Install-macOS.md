# Install macOS

## Step 1: 获取ffmpeg
macOS下推荐使用`brew`来安装`ffmpeg`
- 安装`brew`

    `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

- 安装`ffmpeg`

    `brew install ffmpeg`

## Step 2: 下载xlive

打开[xlive Releases](https://github.com/kira1928/xlive/releases/latest)，选择`xlive-darwin-amd64.7z`下载并解压


## Step 3: 运行

双击 xlive-linux-amd64 执行将使用和程序在同一目录下的 `config.yml` 文件作为默认配置文件启动程序。

也可以执行
```
./xlive-linux-amd64 -c ./config.yml
```
使用指定的配置文件来启动