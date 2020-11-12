package main

import (
	"fmt"
)

func main() {

	// balances multi dim array
	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10},
	}

	fmt.Println(balances)

	nArr := [3]int{1, 2, 3}

	// creates a copy by value, not reference
	mArr := nArr

	fmt.Println("arrays are equal: ", nArr == mArr)

	mArr[0] = 100
	fmt.Println("arrays are equal: ", nArr == mArr)
	fmt.Println("nArr: ", nArr)
	fmt.Println("mArr: ", mArr)

}
