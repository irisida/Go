package main

import "fmt"

func main() {

	// the break statement

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Inside break case. iteration: ", i)
			break
		}
		fmt.Println("for loop iteration: ", i)
	}

	// demo the continue keyword example using the traditional fizzbuzz programming problem

	for j := 0; j < 100; j++ {
		if j > 0 && j%3 == 0 && j%5 == 0 {
			fmt.Println("fizz buzz found at: ", j)
		}
		continue
	}

}
