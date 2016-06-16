package main

import "fmt"

func main() {
	map1 := make(map[string]int, 7)
	map2 := make(map[string]int)
	map3 := map[string]int{}
	map4 := map[string]int{"Mon": 0, "Tue": 1, "Wed": 2, "Thu": 3, "Fri": 4, "Sat": 5, "Sun": 6}
	fmt.Println(map1, map2, map3, map4)
}
