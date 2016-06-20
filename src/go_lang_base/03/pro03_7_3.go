package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var sunny Person
	sunny = Person{"sunny", 41} //最常用的赋值方式，没有写key，必须要按照定义顺序赋值，这种赋值方式不能少写参数
	fmt.Println(sunny)
	sunny = Person{Age: 42, Name: "Go"} //key value赋值，顺序都可以颠倒
	fmt.Println(sunny)
	sunny.Name = "Bill" //这个赋值方式让C程序员很习惯，不过这种赋值方式可以少些某个参数
	sunny.Age = 50
	fmt.Println(sunny)
	sunny.Age, sunny.Name = 10, "Peter" //这个就是具有Go语言特色的赋值方式
	fmt.Println(sunny)
	var peter Person //这里故意不给Name赋值，当然你也可以不给Age赋值。不赋值，系统会给默认值
	peter.Age = 50
	fmt.Printf("Person:%T", peter)
}
