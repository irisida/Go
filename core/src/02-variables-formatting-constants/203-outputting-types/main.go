package main

import "fmt"

/*
We can use the fmt package to output type in different
ways, it's worth highlighting that these different ways
of viewing outputs are exactly that, formats of output
they are not changes of type underneath
*/

func main() {
	var (
		a int
		b byte
	)

	a = 100
	b = 27

	fmt.Printf("type: %8T decimal: %04d binary: %08b hex: %x \n", a, a, a, a)
	fmt.Printf("type: %8T decimal: %04d binary: %08b hex: %x \n", b, b, b, b)

}
