package main

import "fmt"

func main() {
	sunnyMap := map[int]func() string{
		0: func() string { return "aaaa" },
		1: func() string { return "bbbb" },
		2: func() string { return "cccc" },
	}
	fmt.Println(sunnyMap)
}
