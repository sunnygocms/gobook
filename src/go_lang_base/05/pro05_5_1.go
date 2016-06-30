package main

import "fmt"

func main() {
	i := 0
LABEL:
	fmt.Print(i)
	i++
	if i == 10 {
		return
	}
	goto LABEL
}
