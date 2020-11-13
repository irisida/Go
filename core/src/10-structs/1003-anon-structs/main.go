package main

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func main() {

	// anon struct declares the struct and the values directly therafter. Suitable for a
	// case where there's no need for reuse.
	model := struct {
		modelID, manufacturer string
		price                 float64
	}{
		modelID:      "M1",
		manufacturer: "Apple Inc",
		price:        999.00,
	}

	fmt.Printf("%#v\n", model)
	fmt.Printf("Mode Id: %s \n", model.modelID)
}
