package main

import "fmt"

// variadic function that tales 0,1 or more ints
func variadic(x ...int) {
	fmt.Printf("in variadic function: %#v\n", x)
}

func mutatesInput(x ...int) {
	x[0] = 100
	fmt.Printf("in mutatesInput function: %#v\n", x)
}

func main() {

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)

	fmt.Printf("in main - nums is: %#v\n", nums)

	variadic(nums...)

	fmt.Printf("back in main - nums is: %#v\n", nums)

	mutatesInput(nums...)

	fmt.Printf("back in main - nums is: %#v\n", nums)

}
