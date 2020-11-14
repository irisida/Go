package main

import "fmt"

func main() {
	// shortcts for fmt package
	pf := fmt.Printf

	// we use the `&` as the address of operator.

	var x int = 2
	ptr := &x

	pf("address of x  : %p \n", &x)   // using the address of operator
	pf("value of ptr  : %v \n", ptr)  // printing the value
	pf("address of ptr: %p \n", &ptr) // using the address of operator
	pf("type of ptr   : %T \n", ptr)  // printing the value
}
