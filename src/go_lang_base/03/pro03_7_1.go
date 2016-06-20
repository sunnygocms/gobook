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
	string  //string作为匿名字段,当然也可以进一步使用切片slice []string
}

func main() {
	staff := Staff{Person{"Sunny", 41}, "13901234567", "Peking", "5 star staff"}
	fmt.Println(staff.string)
	fmt.Println(staff.Name, staff.Person.Name)
}
