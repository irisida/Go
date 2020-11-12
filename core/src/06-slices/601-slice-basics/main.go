package main

import "fmt"

func main() {
	// declare the slice
	var cities []string

	// check the state and type of a empty slice
	fmt.Println("Cities is nil? : ", cities == nil) // returns true
	fmt.Printf("cities is of type: %#v \n", cities)

	// to append tp a slice we cannot simple use cities[index] = value as this
	// would be out of range right now. instead we can append.

	cities = append(cities, "London")

	fmt.Println("Cities is nil? : ", cities == nil) // should return false now
	fmt.Printf("cities is of type: %#v \n", cities)

	nums := []int{8, 20, 41, 42, 49, 1, 12}
	fmt.Println("Nums: ", nums)
	fmt.Printf("nums: %#v \n", nums)

	beatles := make([]string, 4) // create a slice with 4 empty strings
	fmt.Printf("nums: %#v \n", beatles)

}
