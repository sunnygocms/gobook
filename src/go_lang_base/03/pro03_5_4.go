package main

import "fmt"

func main() {
	slice := []int{5}
	slice = append(slice, 7)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0])
	slice = append(slice, 9)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0])
	s1 := append(slice, 11)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0], "ptr(s1) =", &s1[0])
	s2 := append(slice, 12)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0], "ptr(s2) =", &s2[0])
	fmt.Println("slice:", slice, "s1:", s1, "s2:", s2)
}
