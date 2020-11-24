package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)

	for i := 1; i <= 5; i++ {
		go func(n int, c chan int) {
			f := 1

			for i := 2; i <= n; i++ {
				f *= i
			}
			c <- f
		}(i, ch)

		fmt.Printf("factorial of %d is %d \n", i, <-ch)
	}
}
