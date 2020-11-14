package main

import "fmt"

func main() {

	// creates a pointer to a float64 type
	// is currently uninitialised
	// has a value of nil
	var p1 *float64

	// create a variables that will be inferred as of type float64
	x := 123.45

	// the fact that x is of type float64 allows us to assign
	// the pointer to hold the address of x
	p1 = &x

	// now that p1 holds the address of x
	// we can change x through the dereferencing operator of p1
	*p1 = 999.98

	// lets see the original x - et voila....
	fmt.Println(x) // prints out 999.98

}
