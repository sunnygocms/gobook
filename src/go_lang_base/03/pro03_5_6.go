package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [...]string{"vb", "vc++", "Python", "java", "C lang", "D lang", "Go lang"}
	slice := array[:]
	fmt.Println(slice, "ptr(slice) =", &slice[0])
	sort.Strings(slice)
	//	for _, i := range slice {
	//		fmt.Println(i)
	//	}
	fmt.Println(slice, "ptr(slice) =", &slice[0])
}
