package main

import (
 "fmt"
 "math"
)
type Rectangle struct {
	length int
	breadth int
}
type Circle struct {
	radius float64
}
func (r Rectangle) Area()int {
	return r.length * r.breadth
}
func (c Circle) Area()float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	r := Rectangle {
		length: 4,
		breadth: 3,
	}
	c := Circle {
		radius: 3.5,
	}
	fmt.Println("Area of the Rectangle is ", r.Area())
	fmt.Println("Area of the Circle is ", c.Area())
}
