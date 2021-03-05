package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{
	radius float64
}

type Rectangle struct{
	length,breadth float64
}

func (r *Rectangle) Area() float64 {
	return r.breadth * r.length
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main(){
	var shape Shape
	shape = &Circle{radius: 7}
	fmt.Println(shape.Area())
	shape = &Rectangle{length: 10,breadth: 20}
	fmt.Println(shape.Area())
}