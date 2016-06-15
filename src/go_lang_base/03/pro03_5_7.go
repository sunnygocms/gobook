package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [...]string{"vb", "vc++", "Python", "java", "C lang", "D lang", "Go lang"}
	slice := array[:]
	fmt.Println(slice, "ptr(slice) =", &slice[0])
	sort.Strings(slice) //正序排序
	fmt.Println(slice, "ptr(slice) =", &slice[0])
	sort.Sort(sort.Reverse(sort.StringSlice(slice))) //倒序排序
	fmt.Println(slice, "ptr(slice) =", &slice[0])
}
