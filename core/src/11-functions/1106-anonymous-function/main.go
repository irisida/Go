package main

import "fmt"

func main() {
	func(msg string) { // creates an anonymous function that must be immediately invoked
		fmt.Println(msg)
	}("I am anonymous")

	// similar to the IIFE in javascript
}
