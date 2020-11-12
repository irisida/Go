package main

import (
	"fmt"
	"strings"
)

func main() {
	// declaring an array
	nums := [4]int{70, 80, 90, 100}

	// range across the index, value pairings
	for key, val := range nums {
		fmt.Println(key, val)
	}

	fmt.Println(strings.Repeat("-", 25))

	// loop through with a for loop using the len function as as the controller
	for i := 0; i < len(nums); i++ {
		fmt.Println("index: ", i, "value: ", nums[i])
	}
}
