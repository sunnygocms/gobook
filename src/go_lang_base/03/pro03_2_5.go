package main

import "fmt"

func main() {
	var r1 rune = '\u0041'
	var r2 rune = '\U00000041'
	fmt.Printf("%c  %c\r\n", r1, r2)

	r3 := '中'
	r4 := '\u4E2D'
	r5 := []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
	fmt.Printf("%U  %c  %s\r\n", r3, r4, r5)

	str := []rune(string("养由基")) //string 在下一节详细介绍
	for _, r := range str {
		fmt.Printf("%U\r\n", r)
	}

	fmt.Println(string([]rune{0x517B, 0x7531, 0x57FA}))
	fmt.Println(string([]rune{'\u517B', '\u7531', '\u57FA'}))
}
