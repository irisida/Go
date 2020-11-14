package main

import "fmt"

func one() {
	fmt.Println("I am one")
}

func two() {
	fmt.Println("I am two")
}

func three() {
	fmt.Println("I am three")
}

func main() {
	defer one() // despit being the first called it will be the last executed.
	three()
	two()

	// the defer statement make the action deferred the last one ot take place before
	// exiting the function.
}
