package main

import "fmt"

func main() {
	var (
		a, b int = 101, 142
	)
	c := a & b //按位与
	fmt.Printf("a & b = %4d & %4d ===> %08b  & %08b ====>%4d 换成二进制是：%08b\r\n", a, b, a, b, c, c)

	d := a | b //按位或
	fmt.Printf("a | b = %4d & %4d ===> %08b  | %08b ====>%4d 换成二进制是：%08b\r\n", a, b, a, b, d, d)

	e := a ^ b //按位异或
	fmt.Printf("a | b = %4d & %4d ===> %08b  | %08b ====>%4d 换成二进制是：%08b\r\n", a, b, a, b, e, e)

	f := a &^ b //按位清除，这个比较特殊，就是以后面的数字为准对前面的数字的二进制进行相应的清零操作。
	//如果后面的数字的某个二进制位是1，那么前面数字相应的二进制位就变成0，反之保持原数不变。
	fmt.Printf("a | b = %4d & %4d ===> %08b  | %08b ====>%4d 换成二进制是：%08b\r\n", a, b, a, b, f, f)
}
