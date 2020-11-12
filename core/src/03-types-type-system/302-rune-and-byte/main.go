package main

import "fmt"

func main() {
	var x rune
	var y byte

	fmt.Printf("rune: %T %v \n", x, x) // will show a type of int32 as rune is an alias for int32
	fmt.Printf("byte: %T %v \n", y, y) // will show uint8 as alias for uint8
}
