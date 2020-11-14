package main

import "fmt"

// function with a named return(s)
func namedReturns(x, y int) (a, m int) {
	fmt.Printf("This is namedReturns. I return two values, addition and multiplication of %d and %d \n", x, y)
	a = x + y
	m = x * y
	return
}

func main() {
	v1, v2 := namedReturns(5, 5)
	fmt.Println(v1, v2)
}
