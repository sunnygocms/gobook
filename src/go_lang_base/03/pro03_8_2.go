package main

import (
	"fmt"
	"math"
)

//面积接口
type ShapeArea interface {
	Area() float64 //求面积
}

//周长接口
type ShapePerimeter interface {
	Perimeter() float64 //	求周长
}

//形状接口：包含面积接口、周长接口
type Shape interface {
	ShapeArea
	ShapePerimeter
}

//圆形定义
type Circle struct {
	radius float64
}

//实现圆形面积接口
func (this Circle) Area() float64 {
	return math.Pi * this.radius * this.radius
}

//实现圆形周长接口
func (this Circle) Perimeter() float64 {
	return 2 * math.Pi * this.radius
}

func main() {
	var circle Circle
	circle.radius = 3
	fmt.Println("Cirle area:", circle.Area())
	fmt.Println("Cirle perimeter:", circle.Perimeter())

	//下面是判断是否实现了接口
	var s Shape = circle
	if circleValue, ok := s.(ShapePerimeter); ok {
		fmt.Printf("circle implements Perimeter(): %f\n", circleValue.Perimeter())
	}

}
