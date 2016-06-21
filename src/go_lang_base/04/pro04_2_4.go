package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (this *Person) String() string {
	return fmt.Sprintf("%v", this)
}
func main() {
	sunny := new(Person)
	sunny.Name = "sunny"
	sunny.Age = 41
	fmt.Println(sunny)
}
