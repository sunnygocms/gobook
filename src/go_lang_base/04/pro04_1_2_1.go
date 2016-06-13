package main

import (
	"fmt"
)

var (
	myfun = func(x int, y int) int { //求矩形的面积
		return x * y
	}
)

func myfunc(x, y int) int {
	return func(i, j int) int {
		return i * j
	}(x, y)
}
func main() {
	fmt.Printf("这个矩形的面积是：3 X 5=%d\r\n", myfun(3, 5))
	fmt.Printf("这个矩形的面积是：3 X 5=%d", myfunc(3, 5))
}
