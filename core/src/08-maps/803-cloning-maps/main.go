package main

import "fmt"

func main() {
	src := map[string]float64{"EUR": 0.08, "USD": -0.0677}

	dest := src

	src["EUR"] = 0.1136

	// dest has not been mutated and yet will show the update performed on src
	fmt.Println(dest)

	// maps share the same data structure in memory.
	// to clone, we need to create a new map and then loop over existing map
	// to then copy. An unrelated deep copy is then made.

	newDest := make(map[string]float64)

	for k, v := range src {
		newDest[k] = v
	}

	fmt.Println("After loop, before change: ", src, newDest)

	src["USD"] = -0.202

	fmt.Println("After loop, after change: ", src, newDest)
}
