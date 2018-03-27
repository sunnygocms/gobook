## Linux命令

### 统计代码行数

- 统计当前目录下，go文件数量：

	find . -name "*.go" |wc -l
    
- 统计当前目录下，所有go文件行数：

	find . -name "*.go" |xargs cat|wc -l

- 统计当前目录下，所有py文件行数，并过滤空行：

	find . -name "*.go" |xargs cat|grep -v ^$|wc -l


## 链接
- [目录](https://github.com/sunnygocms/gobook/blob/master/menu.md)