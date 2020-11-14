package main

import "fmt"

func one() {
	defer fmt.Println("I am one") // here we defer the primary action of the function
	three()
	two()
}

func two() {
	fmt.Println("I am two")
}

func three() {
	fmt.Println("I am three")
}

func main() {
	one()
}
