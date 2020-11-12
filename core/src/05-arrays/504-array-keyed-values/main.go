package main

import (
	"fmt"
	"strings"
)

func main() {

	// keyed elements are where we can also specify the index value associated with
	// an element. Important note is that the do not need to be in the correct order
	// only that they are unique, constants and fit within the index range of the
	// defined array size.
	grades := [3]int{
		1: 10,
		0: 5,
		2: 8,
	}

	for key, val := range grades {
		fmt.Println(key, val)
	}

	fmt.Println(strings.Repeat("-", 25))

	// it is not mandatory that we fill all elements
	newGrades := [3]int{
		2: 10,
	}

	for key, val := range newGrades {
		fmt.Println(key, val)
	}
}
