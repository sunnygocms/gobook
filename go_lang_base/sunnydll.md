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

### C语言调用

```c
#include <stdio.h>
#include "libsunny.h"
int main(){
    GoString str;
	GoUint8 show;
    str = Hello();   
    Test();
    printf("%lld\n",str.n);
	show=IsOrientationZero(str);
	printf("%d\n",show);
	return 0;
}
```
###  编译

    gcc main.c -o sunny.exe -I./ -L./ -lsunny
