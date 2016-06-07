package main

import "fmt"

func main() {
	var (
		a, b int     = 3, 2
		c    int     = 8
		x, y float64 = 3.2, 8.0
		z    float64 = 16.0
	)
	fmt.Printf("3 + 2 =%2d\r\n", a+b)
	fmt.Printf("3 - 2 =%2d\r\n", a-b)
	fmt.Printf("2 * 8 =%2d\r\n", b*c)
	fmt.Printf("8 / 2 =%2d\r\n", c/b)
	fmt.Printf("8 %% 2 =%2d\r\n", c%a)
	fmt.Printf("3.2 + 8.0 =%4.2f\r\n", x+y)
	fmt.Printf("3.2 - 8.0 =%4.2f\r\n", x-y)
	fmt.Printf("3.2 * 8.0 =%4.2f\r\n", x*y)
	fmt.Printf("16.0 / 8.0 =%4.2f\r\n", z/y)
	//fmt.Println(z % x) //这个是错误的语句，取模只能是在整数里面

}
