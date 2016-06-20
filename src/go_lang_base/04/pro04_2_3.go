package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (this *Person) String() string {
	return "(名字：" + this.Name + "\t年龄：" + strconv.Itoa(this.Age) + ")\r\n"
}
func main() {
	sunny := new(Person)
	sunny.Name = "sunny"
	sunny.Age = 41
	fmt.Println(sunny)
}
