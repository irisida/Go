package main

import "fmt"

func main() {

	/*
		note that a for is made up of 3 parts:
		1) the index variable of the loop, i
		2) the condition to evaluate against to decide if the loop needs to run
		3) how the index is changed, incremented/decreameted
	*/

	for i := 0; i < 10; i++ {
		fmt.Println("loop value is: ", i)
	}

}
