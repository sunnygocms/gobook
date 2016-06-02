package main

import (
	"flag"
	"fmt"
)

var bFlag = flag.Bool("b", false, "If show message")

func main() {
	//第一个参数，为参数名称，第二个参数为默认值，第三个参数是说明
	name := flag.String("name", "Watson", "Input your name.")
	age := flag.Int("age", 0, "Input your age.")
	flag.Parse()
	if !*bFlag {
		fmt.Println("Mr. Watson, Come Here, I Want You!")
	} else {
		fmt.Println("Hello, ", *name, "   age:", *age)
	}

}
