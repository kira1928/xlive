## Grafana 面板

docker compose 用户可以取消项目根目录下 `docker-compose.yml` 文件中 prometheus 和 grafana 部分的注释以启用统计面板。  
这里是 [设置说明](#%E4%B8%89%E5%90%88%E4%B8%80%E7%AE%80%E5%8C%96%E5%AE%89%E8%A3%85)

非 docker compose 用户需要自行部署 prometheus 和 grafana。  
这里是 [一些建议](#%E6%89%8B%E5%8A%A8%E5%AE%89%E8%A3%85%E7%AC%94%E8%AE%B0)

![image](dashboard.webp)

# 三合一简化安装

## 安装
1. clone repo
>$ git clone https://github.com/kira1928/xlive.git

2. 编辑 .env
自定义账户密码，将example.env保存为.env
默认账户密码 `admin admin`
不改网页里也会提醒你去改。

3. 命令行输入
>$ docker compose up

假如docker 版本低于20，需要安装docker-compose。新的版本是直接内置的

4. 浏览器打开 http://localhost:3000 。

tips
- 使用默认端口和别的端口冲突时，修改相关ports:
- `./Videos:/srv/xlive` 默认保存路径需要自定义


bibliography
1. [Docker-Compose-Prometheus-and-Grafana](https://github.com/Einsteinish/Docker-Compose-Prometheus-and-Grafana) 

# 手动安装笔记
没有 docker 或者想在其他机器配置监控也可以选择手动安装。
虽然API.md里面没有写，但是路径应该是`/api/metrics` 。`prometheus.yml` 需要写成以下形式
``` yml
global:
  scrape_interval: 15s
scrape_configs:
  - job_name: "xlive"
    metrics_path: "/api/metrics"
    scheme: http
    static_configs:
      - targets: ["xlive:8080"] #自行修改ip端口
```
grafana 需要打开浏览器，然后复制[面板内容](/contrib/grafana/dashboard.json)导入

# 群晖（Synology）的情况
[启用 grafana 统计面板](./Synology-related.md#启用-grafana-统计面板)