package main

import "fmt"

func main() {

	// create source slice
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("src type: %T values: %v \n", src, src)

	// take a slice of a slice as a slice...nice!
	dst := src[2:5]
	fmt.Printf("src type: %T values: %v \n", dst, dst)
}
