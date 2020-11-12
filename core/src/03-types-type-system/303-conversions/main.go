package main

import "fmt"

func main() {

	// create a string of a number, returns a rune or unicode character value for the
	// value that was passed, to get the string representation we need Sprintf
	var x = string(99) // comes with a warning

	// Sprintf takes the type and the value
	var y = fmt.Sprintf("%d", 99)
	var z = fmt.Sprintf("%f", 9.98765)

	// printing an int to string is an easy conversion
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
