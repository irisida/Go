package main

import "fmt"

func main() {

	people := [4]string{"Iggy", "Elvis", "Ringo", "Keef"}
	beatles := [4]string{"John", "George", "Paul", "Ringo"}

outer:
	for index, name := range people {
		for _, beatle := range beatles {

			if name == beatle {
				fmt.Printf("Found a beatle, %s at index %d \n", beatle, index)
				break outer
			} else {
				fmt.Printf("No match when comparing %s with %s. \n", name, beatle)
			}
		}
	}

}
