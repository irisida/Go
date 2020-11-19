package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

func print(s shape) {
	fp := fmt.Printf

	fp("Shape: %#v\n", s)
	fp("Area: %v\n", s.area())
	fp("Perimeter: %v\n", s.perimeter())
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// a variable of the type for an interface can hold any
	// underlying type that implements that interface.

	var s shape
	fmt.Printf("type of s: %#v", s)

	ball := circle{radius: 5}
	s = ball
	print(s)
	fmt.Printf("type of s: %#v", s)

	room := rectangle{width: 10, height: 4}
	s = room
	print(s)
	fmt.Printf("type of s: %#v", s)

}
