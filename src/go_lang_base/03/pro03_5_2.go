package main

import "fmt"

func main() {
	slice1 := []byte{'a', 'b', 'c'}
	slice2 := []byte{'d', 'e', 'f'}
	//注意下面这两个区别
	fmt.Println(append(slice1, 'd', 'e', 'f'))
	fmt.Println(append(slice1, slice2...))
}
