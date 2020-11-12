package main

import "fmt"

func main() {
	pt := fmt.Println
	pf := fmt.Printf

	var emps map[string]string
	pt(emps)

	// does not error on querying a non existent element
	// gracefully returns
	pf("Value for %q is %q \n", "programmer", emps["programmer"])

	// the comma-ok idiom
	_, ok := emps["Ed"]

	if !ok {
		pt("Ed was not found in the map")
	}

	// generally maps in Go were design for fast lookup rather than fast looping so
	// we can use a standard for loop construct to traverse over a map

	bals := map[string]float64{
		"USD": 0.022,
		"GBP": -0.003,
		"EUR": 0.11,
	}

	// use the shortcut for fmt.Println
	for k, v := range bals {
		pt(k, v)
	}
}
