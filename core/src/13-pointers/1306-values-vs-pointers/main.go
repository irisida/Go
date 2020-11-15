package main

import "fmt"

// changes are made to the value of the underlying memory segment
// therefore the changes will persist after the function returns.
func changesTheValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	pf := fmt.Printf

	pf("\tPointer function 01 -> quantity: %d price: %.2f name: %s Sold: %v\n", *quantity, *price, *name, *sold)

	*quantity = 100
	*price = 1199.99
	*name = "Macbook Air 2020-M1"
	*sold = true

	pf("\tPointer function 02 -> quantity: %d price: %.2f name: %s Sold: %v\n", *quantity, *price, *name, *sold)
}

// changes made here will be localised.
func changesTheValues(quantity int, price float64, name string, sold bool) {
	pf := fmt.Printf

	pf("\tValue function 01 -> quantity: %d price: %.2f name: %s Sold: %v\n", quantity, price, name, sold)

	quantity = 100
	price = 1299.99
	name = "Macbook Air 2020-M1"
	sold = true

	pf("\tValue function 02 -> quantity: %d price: %.2f name: %s Sold: %v\n", quantity, price, name, sold)
}

func main() {

	pf := fmt.Printf

	q := 10
	p := 1799.50
	n := "MBA 2020 model"
	s := false

	pf("in main 01 -> quantity: %d price: %.2f name: %s Sold: %v\n", q, p, n, s)

	changesTheValues(q, p, n, s)

	pf("in main 02 -> quantity: %d price: %.2f name: %s Sold: %v\n", q, p, n, s)

	changesTheValuesByPointer(&q, &p, &n, &s)

	pf("in main 03 -> quantity: %d price: %.2f name: %s Sold: %v\n", q, p, n, s)
}
