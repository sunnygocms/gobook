package main

import "fmt"

func main() {
	var a, b, c, d byte
	a = 'A' //在 ASCII 码表中，A 的值是 65，而使用 16 进制表示则为 41,下面四种写法是等效的
	b = 65
	c = '\x41'
	d = 0x41
	fmt.Println(a, b, c, d)
	fmt.Printf("%c  %c  %c  %c", a, b, c, d)
}
