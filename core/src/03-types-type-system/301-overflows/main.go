package main

import "fmt"

func main() {

	/*
		you cannot initially assign a value that overflows the type, this will
		generate a compile time error, but one that overflows in runtime will
		be handled for you gracefully by boomeranging to the first value at the
		opposite end of the value range for that type.
	*/

	var i8 int8 = 127
	var ui8 uint8 = 255

	fmt.Printf("i8: %v  ui8: %v \n", i8, ui8)

	// blow the value ranges of types
	i8++
	ui8++

	// show new output
	fmt.Printf("i8: %d  ui8: %d \n", i8, ui8)
}
