package main

import (
	"fmt"
)

//全局变量
var (
	length float64 = 3
	width  float64 = 5
	area   float64 = 15
)

func Area(length float64, width float64) float64 {
	fmt.Println("形式参数(length):", length, "形式参数(width):", width)
	return length * width
}

func main() {
	fmt.Println("全局变量(length):", length, "全局变量(width):", width)
	fmt.Println("全局变量(area):", area)
	area := Area(7, 8) //注意此处我传入的length和width与全局变量的值不同
	fmt.Println("局部变量(area):", area)
}
