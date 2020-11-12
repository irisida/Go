package main

import "fmt"

/*
We're going to see static typing in effect here in this sample
program. We'll use simple variable types to demonstrate the
change of values within type but that type leaping isn't an
option in Go.
*/

func main() {
	var a = 4
	var b = 9.9

	fmt.Printf("%T %v \n", a, a) // returns int 4
	fmt.Printf("%T %v \n", b, b) // returns float64 9.9

	// a = b // this would generate a compile error
	// b = a // this would generate a compile error also

	a = 100
	b = 100

	fmt.Printf("%T %v \n", a, a) // returns int 100
	fmt.Printf("%T %v \n", b, b) // returns float64 100

}
