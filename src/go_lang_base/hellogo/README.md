##生成有图标的windows执行程序
```bat
go get github.com/akavel/rsrc
cd %GOPATH%/src/github.com/akavel/rsrc
go build
copy rsrc.exe %GOROOT%\bin\.
```
  
现在创建一个目录，比如起名为hellogo

在这个目录下存入一个ico文件，比如叫favicon.ico

在这个目录下执行：
rsrc -ico favicon.ico -o my.syso

 

随便写一个go程序，比如叫做sunnny.go,为什么起的名字和目录的名字不同哪？是为了告诉你不能够使用 go build sunny.go来编译

而要使用 go build来编译

 

编译完成，生成了有图标的windows hellogo.exe程序。