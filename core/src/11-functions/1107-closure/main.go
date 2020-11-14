package main

import "fmt"

func inc(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	a := inc(10)

	for i := 0; i < 10; i++ {
		fmt.Printf("in inc loop index: %d increment value: %d \n", i, a())
	}

}
