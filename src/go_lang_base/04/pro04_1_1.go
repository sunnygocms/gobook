package main

import (
	"fmt"

	"./sunny"
)

func main() {
	sunny.Fun1()                                //无参数，无返回值函数调用
	sunny.Fun2("function", "Fun2")              //多个参数，没有返回值函数调用
	fmt.Println(sunny.Fun3("function", "Fun3")) //多个参数，一个返回值，不指定返回值名称函数调用
	fmt.Println(sunny.Fun4("function", "Fun4")) //多个参数，一个返回值，指定返回值名称函数调用
	msg, err := sunny.Fun5("function", "Fun5")  //多个参数,还包括可变长参数；多个返回值函数调用，不过没有传递可变长参数

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	msg, err = sunny.Fun5("function", "Fun5", "a") //多个参数,还包括可变长参数；多个返回值函数调用，且传递了可变长参数
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
	msg, err = sunny.Fun5("function", "Fun5", "a", "b", "c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}
