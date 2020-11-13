package main

import "fmt"

func f1() {
	fmt.Println("This is f1. It takes no parameters")
}

func f2(x, y int) {
	fmt.Printf("This is f2. It takes two parameters x: %d y: %d \n", x, y)
}

func f3(x, y int) int {
	fmt.Printf("This is f3. i was passed x: %d y: %d I am returning x+y = %d \n", x, y, x+y)
	return x + y
}

func f4(x, y int) (int, int) {
	fmt.Printf("This is f4. I return two values, addition and multiplication of %d and %d \n", x, y)
	return x + y, x * y
}

func main() {
	// call a basic programmer defined function
	f1()

	// call a function with parameters
	f2(10, 20)

	// function with a return
	v := f3(5, 5)
	fmt.Println("after f3 ran the value is: ", v)

	v1, v2 := f4(5, 5)
	fmt.Printf("after f4 ran the values are: %d and %d \n", v1, v2)
}
