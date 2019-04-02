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

//export Summ
func Summ(x, y int) int {
	return x + y
}

//export Hello
func Hello() string {
	return "1.jpg"
}

//export Test
func Test() {
	println("export Test")
}
