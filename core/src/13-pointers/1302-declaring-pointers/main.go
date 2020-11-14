package main

import "fmt"

func main() {
	// shortcts for fmt package
	pf := fmt.Printf

	var ptr1 *float64 // creates an uninitalised pointer to a float64 type, value of nil
	pf("ptr1 - Type: %T Address: %p Value: %v\n", ptr1, &ptr1, ptr1)

	// create a float64 and assign a value, assign the ptr1 to hold address of this new value
	dbl := 123.45
	ptr1 = &dbl
	pf("ptr1 - Type: %T Address: %p Value: %v\n", ptr1, &ptr1, *ptr1)

	// using the new() function
	ptr2 := new(int)
	pf("ptr2 - Type: %T Address: %p Value: %v\n", ptr2, &ptr2, *ptr2)

	// create a new int and assign a value, point ptr2 to the address of thenew variable
	val := 100
	ptr2 = &val
	pf("ptr2 - Type: %T Address: %p Value: %v\n", ptr2, &ptr2, *ptr2)

	// update the underlying values through the pointer dereferencing capabilities
	*ptr1 = 99.99
	*ptr2 = 500

	// show updates
	pf("ptr1 - Type: %T Address: %p Value: %v\n", ptr1, &ptr1, ptr1)
	pf("ptr2 - Type: %T Address: %p Value: %v\n", ptr2, &ptr2, *ptr2)

	// verify with the source variables
	pf("dbl  - Type: %T Address: %p Value: %v\n", dbl, &dbl, dbl)
	pf("val  - Type: %T Address: %p Value: %v\n", val, &val, val)
}
