package main

import "fmt"

// Product comment
type Product struct {
	price float64
	name  string
}

// mutation by value demonstrates that for structs too Go will create a
// local copy and the changes are applicable to that copy and the scope
// of the function only.
func changeProduct(p Product) {
	pf := fmt.Printf

	pf("\tchangeProduct function 01 -> Product: %v\n", p)

	p.price = 319.49
	p.name = "updated-ProductName"

	pf("\tchangeProduct function 02 -> Product: %v\n", p)
}

// same as above but with pointers
func changeProductWithpointer(p *Product) {
	pf := fmt.Printf

	pf("\tchangeProduct function 01 -> Product: %v\n", *p)

	p.price = 529.99
	p.name = "ChangedInPointerFunction"

	pf("\tchangeProduct function 02 -> Product: %v\n", *p)
}

func main() {
	pf := fmt.Printf

	prod := Product{
		price: 1299.99,
		name:  "macbook air M1",
	}

	pf("in Main 01 Product:  %v\n", prod)

	changeProduct(prod)

	pf("in Main 02 Product:  %v\n", prod)

	changeProductWithpointer(&prod)

	pf("in Main 03 Product:  %v\n", prod)
}
