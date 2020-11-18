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
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}

	print(c1)
	print(r1)
}
