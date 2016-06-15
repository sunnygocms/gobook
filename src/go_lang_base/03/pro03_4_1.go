package main

import "fmt"

/**
 * [100]float32
 * [2][3]int
 * [3*10]byte //GO定义数组的时候可以采用表达式作为参数，表达式只能是由常量组成，不能是变量
 * [3][2][2]float32  // 和 [3]([2]([2]float32)) 是一个意思
 */
func main() {
	const N = 2
	x := [2 * N]int{1, 2, 3, 4}
	for _, i := range x {
		fmt.Println(i)
	}
	array := [...]string{"vb", "vc++", "Python", "java", "C lang", "D lang", "Go lang"}
	for _, i := range array {
		fmt.Println(i)
	}
	//数组长度计算
	fmt.Println(len(array))
}
