package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr = []int{8, 2, 1, 0, 3}                     //手机号码只是由这5个数字组成
	var index = []int{2, 0, 3, 2, 4, 0, 1, 3, 2, 3, 3} //手机号码各位的顺序
	tel := ""
	for _, i := range index {
		tel += strconv.Itoa(arr[i])
	}
	fmt.Println("联系方式：", tel)
}
