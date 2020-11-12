package main

import "fmt"

func main() {

	/*
		There is no while loop in Go, however we can perform the same functionality with a for loop that is
		coopted syntactically to how we need it to be
	*/

	conditionalValue := 20

	// for and specify the closing/stop condition
	for conditionalValue > 0 {
		fmt.Println("While loop in go: ", conditionalValue)

		// ensure we interact with the controller
		conditionalValue--
	}

}
