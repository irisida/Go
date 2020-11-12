package main

import "fmt"

type age int

func main() {

	var i1 = 21
	var a1 age = 21

	// display the values and types
	fmt.Printf("%T, %v \n", i1, i1)
	fmt.Printf("%T, %v \n", a1, a1)

	// perform a comparison
	var b1 bool
	//b1 = i1 == a1  // will cause a compile time error because the types do not match.
	b1 = i1 == int(a1) // requires an explicit conversion to be able to perform a value equality check
	fmt.Println(b1)
}
