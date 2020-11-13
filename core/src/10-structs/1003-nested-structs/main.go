package main

import "fmt"

type address struct {
	houseNumber                  int
	street, town, city, postcode string
}

// composition of the structure contains the address structure
type person struct {
	fname, lname string
	address      address
}

func main() {

	p1 := person{}

	p1.fname = "John"
	p1.lname = "Lennon"
	p1.address.houseNumber = 123 // note the dot notation for accessing nested structs
	p1.address.street = "Liverpool Road"
	p1.address.city = "Liverpool"
	p1.address.postcode = "LI822P"

	fmt.Println(p1)

}
