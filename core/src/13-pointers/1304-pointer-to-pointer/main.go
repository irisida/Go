package main

import "fmt"

func main() {

	p := fmt.Printf

	val := 100.0
	p1 := &val

	pp1 := &p1

	p("Type: %T Address: %p Value: %v\n", val, &val, val)
	p("Type: %T Address: %p Value: %v\n", p1, &p1, *p1)
	p("Type: %T Address: %p Value: %v\n", pp1, &pp1, *pp1)

	**pp1 = 500.

	p("Type: %T Address: %p Value: %v\n", val, &val, val)
	p("Type: %T Address: %p Value: %v\n", p1, &p1, *p1)
	p("Type: %T Address: %p Value: %v\n", pp1, &pp1, *pp1)
}
