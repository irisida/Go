package main

import "fmt"

/*
We're going to see static what are called empty or zero
values. in Go variables declared and not assigned a value
take the default or zero value for the underlying type.
An uninitialised variable where the type is not specified
will cause a compilation error.
*/

func main() {
	var (
		a int
		b float64
		c string
		d byte
		e bool
	)

	fmt.Println(a) // 0
	fmt.Println(b) // 0
	fmt.Println(c) // note that string prints "" essentially it prints an empty string.
	fmt.Println(d) // 0
	fmt.Println(e) // defaults to false

}
