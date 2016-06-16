package main

import "fmt"

func main() {
	sunnyMap := map[string]string{"Mon": "一", "Tue": "二", "Wed": "三", "Thu": "四", "Fri": "五", "Sat": "六", "Sun": "日"}
	delete(sunnyMap, "Mon")
	fmt.Println(sunnyMap)
}
