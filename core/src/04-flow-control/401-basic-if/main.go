package main

import "fmt"

// let's create a simple struct that hold two fields
type model struct {
	modelName string
	price     float64
}

func (m model) getPrice() float64 {
	return m.price
}

func (m model) getModelName() string {
	return m.modelName
}

func main() {

	// create the objects
	mac1 := model{"mbp", 2299.00}
	mac2 := model{"air", 1649.00}

	// set our value to use in the conditional checks
	maxBudget := 2800.00

	// flow controller
	if mac1.getPrice() <= maxBudget {
		fmt.Printf("%s starts at %.2f \n", mac1.getModelName(), mac1.getPrice())
	} else if mac2.getPrice() <= maxBudget {
		fmt.Printf("%s costs %.2f \n", mac2.getModelName(), mac2.getPrice())
	} else {
		fmt.Printf("No models inside budget at this time")
	}

}
