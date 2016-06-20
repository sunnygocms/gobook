package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Staff struct {
	Person  //匿名字段
	Phone   string
	Address string
	Name    string //重名覆盖
}
type StructA struct {
	Name string
}
type StructB struct {
	Name string
	Age  int
}
type StructC struct {
	StructA
	StructB
}

func main() {
	staff := Staff{Person{"Sunny", 41}, "13901234567", "Peking", "Knight"}
	fmt.Println(staff.Name, staff.Person.Name)
	s := StructC{StructA{"张三"}, StructB{"李四", 30}}
	fmt.Println(s.Name)
}
