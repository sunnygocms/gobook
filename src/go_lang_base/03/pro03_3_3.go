package main

import "fmt"

func main() {

	fmt.Printf(" 8 << 2 = %2d\r\n", 8<<2)
	fmt.Printf(" 2 << 5 = %2d\r\n", 2<<5)
	fmt.Printf("64 >> 3 = %2d\r\n", 64>>3)
	fmt.Printf("64 >> 6 = %2d\r\n", 64>>6)
	fmt.Printf("64 >> 7 = %2d\r\n", 64>>7) //64换算成二进制是01000000，右位移7位就变成0
}
