package main

import (
	"fmt"
	"sort"
)

func main() {
	sunnyMap := make(map[int]string)
	sunnyMap[0] = "Mon"
	sunnyMap[1] = "Tue"
	sunnyMap[2] = "Wed"
	sunnyMap[3] = "Thu"
	sunnyMap[4] = "Fri"
	sunnyMap[5] = "Sat"
	sunnyMap[6] = "Sun"

	var keys []int
	for key, _ := range sunnyMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, i := range keys {
		fmt.Printf("%d->%s\r\n", i, sunnyMap[i])
	}
}
