package main

import (
	"fmt"
	"math/rand"
)

func Yourfun() int {
	return rand.Intn(10)
}
func main() {
	switch i := Yourfun(); {
	case i < 3:
		fmt.Println("small", i)
	case i > 5:
		fmt.Println("big", i)
	default:
		fmt.Println("Normal")
	}
}
