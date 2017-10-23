package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	chs := make([]chan int, 4)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int)
		go func(i int) {
			for {
				chs[i] <- i + 1
			}
		}(i)
	}
	f := make([]bytes.Buffer, len(chs))

	for i := 0; i < 10; i++ {
		for j := 0; j < len(f); j++ {
			fmt.Fprintf(&f[j], "%d ", <-chs[(i+j)%len(chs)])
		}
	}
	for i := 0; i < len(f); i++ {
		fmt.Printf("%s: %s\n", string(65+i), f[i].String())
		file, _ := os.Create(`./` + string(65+i) + `.txt`) //创建文件
		file.Write([]byte(f[i].String()))
		file.Close()
	}
}
