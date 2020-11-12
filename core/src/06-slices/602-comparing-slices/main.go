package main

import "fmt"

func main() {

	var n []int
	fmt.Println(n == nil)

	m := []int{} // empty but it is initialised
	fmt.Println(m == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// fmt.Println(a == b) - compilation error, slices comparison only works with nil

	a = nil
	var eq bool = true
	for i, val := range a {
		if val != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a & b are equal")
	} else {
		fmt.Println("a & b are not equal")
	}

}
