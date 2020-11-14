package main

import "fmt"

// variadic function that tales 0,1 or more ints
func variadic(x ...int) {
	fmt.Printf("%T\n", x)
	fmt.Printf("%#v\n", x)
}

func main() {
	// call with zero params
	variadic()

	// single param
	variadic(1)

	// multiple params
	variadic(1, 2, 3)

	// a good example of variadic functions in the standard library is the append function
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	variadic(nums...)
}
