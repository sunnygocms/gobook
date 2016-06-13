package main

import (
	"fmt"
)

type Int int

func (t *Int) myadd(i, j int) int {

	return i + j

}
func main() {
	var t Int
	fmt.Println(t.myadd(1, 2))
}
