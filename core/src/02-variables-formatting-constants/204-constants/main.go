package main

import "fmt"

/*
All basic literals are constants, 5 will always be 5 just
as the string "Hello, World" will always be the same. We
can therefore say that all literal are unnamed constants.
In Go a constant belongs to compile-time, its value is
read and understood, it is not allowed to change during
the running of the program.
*/

func main() {
	const (
		a int     = 500
		b float64 = 41.99
	)

	// a = 100		// compile error
	// b = 27.99	// compile error

	fmt.Printf("type: %8T decimal: %d  \n", a, a)
	fmt.Printf("type: %8T decimal: %.02f \n", b, b)

}
