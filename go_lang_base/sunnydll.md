# go 生成so 文件方法
### 一直以来不会使用C++一直是我心头痛，不过学习Go从某种意义上补偿了我这个遗憾。
### 比如生成dll一直以来几乎就是C和C++的专利，现在我可以用Go轻松的实现这一点。

下面我使用一个实例来做到这一点,这个是我写的判断jpeg图是不是正向的一个Go的函数

``` go
package main

import "C"
import (
	"fmt"
	"os"

	SunnyUtil "github.com/sunnyregion/util"
)

func main() {}

// IsOrientationZero  是否是正向的照片
// @Param fname string  文件名
// @Return b bool 返回1是true
//export IsOrientationZero
func IsOrientationZero(fname string) (b bool) {
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	if b, err = SunnyUtil.IsOrientationZero(f); err != nil {
		b = false
	}
	return b
}
```
### 下面编译

    go build -x -v -ldflags "-s -w" -buildmode=c-shared -o libsunny.so main.go

### 生成 libsunny.h和libsunny.so
### 特别注意  *** //export IsOrientationZero *** 没有这个export，有可能无法生成的头文件(.h)，还不会有任何报错。
### C语言调用

```c
#include <stdio.h>
#include "libsunny.h"
int main(){
  GoUint8 show;
  GoString name = {"1.jpg", 5};
  show=IsOrientationZero(name);
  printf("%d\n",show);

	// 横着的
	GoString name2 = {"2.jpg", 5};
	show=IsOrientationZero(name2);
	printf("%d\n",show);
	return 0;
  return 0;
}
```
###  编译

    gcc main.c -o sunny.exe -I./ -L./ -lsunny


### 在ubuntu下需要 把libsunny.so拷贝到/usr/lib中用于运行。在Mac下没有这个问题。

### 解决方法

### -添加环境变量。
	
	LD_LIBRARY_PATH=.
	export LD_LIBRARY_PATH

### -编译的时候添加参数

	gcc main.c -o sunny.exe -I./ -L./ -lsunny -Wl,-rpath=.
	
### -写成makefile文件

```Makefile
#g++ compiler
CC = gcc
CCFLAGS = -I. -Wall -fmessage-length=0 -fPIC
LDFLAGS = -Wl,-rpath,'./'

#Debug or Release, Debug:-g Release:-O3
ifdef GDB
	OPTS = -g -rdynamic
else
	OPTS = -O2 -rdynamic
endif

#Link the .so or .a library
LIBS = -L./ -lsunny

#Target file
TARGET = sunny.exe

#Gernation the target file
all:$(TARGET) 
.PHONY : clean

objects = main.o

$(TARGET):$(objects)
	$(CC) -o $@ $(objects) $(LIBS) $(LDFLAGS)

%.o : %.c
	$(CC) $(CFLAGS) $(OPTS) -c $<

clean:
	-rm -f $(objects) sunny.exe

```

### 检查
	
	ldd sunny.exe
	
### 看到 

	libsunny.so => ./libsunny.so (0x00007fccc8751000)
	
### 就ok。