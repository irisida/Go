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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {

	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	// calling s.volume() will be an error
	// using type assertion allows us to access type
	// dynamic functionalities.

	ball, ok := s.(circle)
	if ok {
		fmt.Printf("Volume: %v \n", ball.volume())
	}

	// change to rectange type tp trigger the case below
	// comment out for main circle flow of the example.
	s = rectangle{width: 2, height: 5}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type", value)

	case rectangle:
		fmt.Printf("%#v has rectangle type", value)
	}
}
