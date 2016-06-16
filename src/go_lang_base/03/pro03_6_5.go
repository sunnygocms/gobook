package main

import "fmt"

func main() {
	sunnyMap := map[string]string{"Mon": "一", "Tue": "二", "Wed": "三", "Thu": "四", "Fri": "五", "Sat": "六", "Sun": "日"}
	val1, isExist1 := sunnyMap["Sat"]
	val2, isExist2 := sunnyMap["sat"]

	fmt.Println("Sat is exist?", isExist1, "value:", val1)
	fmt.Println("sat is exist?", isExist2, "value:", val2)
}
