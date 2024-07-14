# Install-Linux

## Step 1: 获取ffmpeg
- Ubuntu / Debian 等

    `sudo apt-get install ffmpeg`

- Fedora / CentOS 等

    `sudo yum install ffmpeg`

- ArchLinux

    `sudo pacman -S ffmpeg`

- Build from source code
    
    [FFmpeg Compilation Guide](https://trac.ffmpeg.org/wiki/CompilationGuide)

## Step 2: 下载xlive

打开[xlive Releases](https://github.com/kira1928/xlive/releases/latest)，下载对应的版本并解压

## Step 3: 注册为服务（可选）
- systemd  
注意替换`ExecStart`中的二进制文件和配置文件路径
```shell
echo "[Unit]
Description=Live Stream Saver
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
ExecStart=/usr/bin/xlive -c /etc/xlive/config.yml
Restart=on-failure

[Install]
WantedBy=multi-user.target" > /usr/lib/systemd/system/xlive.service

systemctl daemon-reload

systemctl enable xlive.service

systemctl start xlive.service
```
