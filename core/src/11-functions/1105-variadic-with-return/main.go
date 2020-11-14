package main

import "fmt"

func sumAndProduct(x ...int) (s int, p int) {
	s = 0
	p = 1

	for _, v := range x {
		s += v
		p *= v
	}

	return
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	sum, prod := sumAndProduct(nums...)
	fmt.Printf("sum: %d  product: %d", sum, prod)
}
