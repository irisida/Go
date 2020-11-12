package main

import "fmt"

type age int

func (a age) isAnEligibleVoter() bool {
	if a >= 18 {
		return true
	}
	return false
}

func main() {
	var p1 age = 21
	fmt.Println(p1.isAnEligibleVoter())

	var p2 age = 17
	fmt.Println(p2.isAnEligibleVoter())
}
