package main

import "fmt"

func main() {
	slice := []int{5}
	slice = append(slice, 7)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0])
	slice = append(slice, 9)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0])
	s1 := make([]int, len(slice))
	copy(s1, slice)
	s1 = append(s1, 11)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0], "ptr(s1) =", &s1[0])
	s2 := make([]int, len(slice))
	copy(s2, slice)
	s2 = append(s2, 12)
	fmt.Println("cap(slice) =", cap(slice), "len(slice) =", len(slice), "ptr(slice) =", &slice[0], "ptr(s2) =", &s2[0])
	fmt.Println("slice:", slice, "s1:", s1, "s2:", s2)
}
