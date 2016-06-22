package main

import (
	"fmt"
	"math"
)

//接口定义
type Shape interface {
	Area() float64 //求面积
}

//圆形定义
type Circle struct {
	radius float64
}

//圆形面积计算
func (this Circle) Area() float64 {
	return math.Pi * this.radius * this.radius
}

//矩形定义
type Rectangle struct {
	length float64
	width  float64
}

//矩形面积计算
func (this Rectangle) Area() float64 {
	return this.length * this.width
}
func main() {
	var circle Circle
	circle.radius = 3
	fmt.Println(circle.Area())

	rec := Rectangle{5.0, 8}
	fmt.Println(rec.Area())
}
