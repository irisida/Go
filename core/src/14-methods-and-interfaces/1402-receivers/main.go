package main

import (
	"fmt"
)

type names []string

func (n names) print() {
	// makes itself an available method on variables of type names. it is called using the dot operator

	p := fmt.Println

	for i, name := range n {
		p(i, name)
	}
}

func main() {

	beatles := names{"Ringo", "George", "Paul", "John"}

	beatles.print()

}
