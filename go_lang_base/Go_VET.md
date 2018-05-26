## Go 代码检查

### 统计代码行数

推荐我开发的开源软件--[artHummer](https://github.com/sunnyregion/artHammer)，目前只是简单的统计代码，后续可能开发更加复杂的功能。

### go report
如果您的代码是放在github上面的，可以登录[go report](https://goreportcard.com/)  进行代码扫描。

### gofmt 整理代码

```go
	gofmt  -l -w -s *.go
 	gofmt -r '(a)->a' -l -w *.go
```

### go tool vet
这个是go语言自带的检测工具，检测的是一些语法错误，比如执行不到的返回，错误的调用，defer执行的之前没有判断err什么的。

不过这个工具最好的地方时能够检测struct里面注释的错误。看一个我遇到的[例子](https://github.com/astaxie/beego/issues/3160)

### golint检测
源码下载[地址](github.com/golang/lint/golint)

		go get  github.com/golang/lint/golint

这个工具检测非常严格，比如：

- package xxx 没有写在第一行
- func 、struct 没有写说明
 - err返回没有放在最后一个
 - 变量和函数名包含以下单词必须全部大写：
     - url ==> URL
     - json ==> JSON
     - html ==> HTML
     - id ==> ID



## 链接
- [目录](https://github.com/sunnygocms/gobook/blob/master/menu.md)