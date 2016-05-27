package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// @Description show person's message
// @Param	name  string
// @Param	age  int

func (student *Person) show(name string, age int) {
	student.Name = name
	student.Age = age
	fmt.Printf("name: %s , Age: %d\r\n", student.Name, student.Age)
}

func main() {
	fmt.Println("Mr. Watson, Come Here, I Want You! \r\n 沃特森先生,过来，我想见你！") //March 10, 1876: ‘Mr. Watson, Come Here … ‘
	var p Person
	p.show("Bell", 42)
}
