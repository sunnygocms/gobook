package main

import "fmt"

func main() {
	sunnyMap := map[string][]string{"a": {"一", "a", "b"}, "b": {"二", "hello"}, "c": {"三"}}
	sunnyMap["c"] = append(sunnyMap["c"], "hhh")
	fmt.Println(sunnyMap)
}
