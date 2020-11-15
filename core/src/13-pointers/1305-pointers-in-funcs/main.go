package main

import "fmt"

// Go passes by value so what is passed is localised
func changeWithPointer(a *int) *float64 {
	pf := fmt.Printf

	pf("\tInside changeWithPointer, a = %d when received\n", *a)
	*a = 100
	pf("\tInside changeWithPointer, a = %d after update\n", *a)

	b := 999.98
	return &b
}

func changeWithValue(a int) float64 {
	pf := fmt.Printf

	pf("\tInside changeWithValue, a = %d when received\n", a)
	a = 100
	pf("\tInside changeWithValue, a = %d after update\n", a)

	b := 999.98
	return b
}

func main() {
	pf := fmt.Printf

	x := 8

	pf("With value before change x = %d\n", x)
	changeWithValue(x)
	pf("With value after change x = %d\n", x)

	// create a pointer
	p := &x
	pf("With pointer before change x = %d\n", x)
	changeWithPointer(p)
	pf("With pointer after change x = %d\n", x)
}
