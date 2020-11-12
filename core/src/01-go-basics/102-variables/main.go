package main

import (
	"fmt"
)

/*
1. A variable is a name for a memory location where a specific value is stored.
2. Variables belong to and are created in runtime.
3. A declared variables must be used or removed. Unused variables will cause an error
4. _ is used as the blank identifier. This can be used to mask unused variables
*/
func main() {
	// work with simple variables
	simpleVariables()

	// multiple variables declarations
	multipleVariables()

	// group declaration style
	groupedVariables()
}

func simpleVariables() {
	// using long form declaration
	var num1 int = 30

	// using short declaration operator
	num2 := 33

	fmt.Printf("type: %T value: %v \n", num1, num1)
	fmt.Printf("type: %T value: %v \n", num2, num2)
}

func multipleVariables() {
	// style 1
	var name string
	// note that multiple variables of the same type can have a
	// short circuited declaration to avoid writing int everywhere
	var age, score1, score2, score3 int

	name = "student"
	age = 25
	score1 = 70
	score2 = 82
	score3 = 79

	fmt.Printf("%s is %d years old. They scored %v, %v & %v over the three tests \n", name, age, score1, score2, score3)
}

func groupedVariables() {
	var (
		productName string
		modelID     string
		price       float64
	)

	productName = "newProduct of some kind"
	modelID = "ABCeasyAs123"
	price = 129.99

	fmt.Printf("%s, %s costs £%.2f \n", productName, modelID, price)
}
