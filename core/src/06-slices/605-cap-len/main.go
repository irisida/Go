package main

import "fmt"

func main() {

	var src []int
	fmt.Printf("len: %v  cap: %v\n", len(src), cap(src))

	src = append(src, 1, 2)
	fmt.Printf("len: %v  cap: %v\n", len(src), cap(src))

	src = append(src, 3)
	fmt.Printf("len: %v  cap: %v\n", len(src), cap(src))

	src = append(src, 4)
	fmt.Printf("len: %v  cap: %v\n", len(src), cap(src))

	src = append(src, 5)
	fmt.Printf("len: %v  cap: %v\n", len(src), cap(src))
}
