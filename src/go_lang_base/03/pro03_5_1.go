package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 4)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", slice1)
	}
}
