package main

import "fmt"

func main() {
	i := 0

	// this works like a for loop. Not the only variables operated on are already
	// declared and that no new variables are allowed in a label/goto block
loop:
	if i < 10 {
		fmt.Println("i is: ", i)
		i++
		goto loop
	}
}
