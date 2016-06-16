package main

import "fmt"

func main() {
	sunnyMap := make(map[int]string)
	sunnyMap[0] = "Mon"
	sunnyMap[1] = "Tue"
	sunnyMap[2] = "Wed"
	sunnyMap[3] = "Thu"
	sunnyMap[4] = "Fri"
	sunnyMap[5] = "Sat"
	sunnyMap[6] = "Sun"

	for key, value := range sunnyMap {
		fmt.Printf("%d->%s\r\n", key, value)
	}
}
