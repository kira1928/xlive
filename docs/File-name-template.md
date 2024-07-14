
## 输出文件名模板

默认的文件名模板如下
```
{{ .Live.GetPlatformCNName }}/{{ .HostName | filenameFilter }}/[{{ now | date "2006-01-02 15-04-05"}}][{{ .HostName | filenameFilter }}][{{ .RoomName | filenameFilter }}].flv
```
![](./filename-tmpl.svg)

bgo最终输出的文件名为`${out_put_path}/${out_put_tmpl}`。  
`out_put_tmpl`基于[go template](https://golang.org/pkg/text/template/)实现，并添加[sprig](http://masterminds.github.io/sprig/)中的方法和如下的方法：

| 方法 | 说明 |
| -- | -- |
| decodeUnicode | 还原Unicode字符串，比如`\\u9ed1\\u6697\\u5251` -> `黑暗剑` |
| replaceIllegalChar | 替换文件名中不支持的字符为`_`，包含：`/`, `\`, `:`, `*`, `?`, `"`, `<`, `>`, `\|` |
| unescapeHTMLEntity | 同[html.UnescapeHTMLEntity](https://golang.org/pkg/html/#UnescapeString)，处理html转义 |
| filenameFilter | `replaceIllegalChar` + `unescapeHTMLEntity` |

输入的对象为[`live.Info`](https://github.com/kira1928/xlive/blob/master/src/live/info.go#L7)，可以通过如下指令拿到相应的信息

| 指令 | 说明 |
| -- | -- |
| `.HostName` | 主播名 |
| `.RoomName` | 房间名 |
| `.Status`, `.Listening`, `.Recoding` | 直播状态，bgo监控器状态，bgo录制器状态，永远为`True` |
| `.Live.GetLiveId` | 直播间ID（一串md5 string，由bgo根据直播间url计算所得）|
| `.Live.GetRawUrl` | 直播间url |
| `.Live.GetPlatformCNName` | 直播平台中文名 |

