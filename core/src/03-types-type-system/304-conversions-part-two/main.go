package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "255"
	s2 := "500"
	s3 := "99.99"

	fmt.Printf("%T %v \n", s1, s1)
	fmt.Printf("%T %v \n", s2, s2)
	fmt.Printf("%T %v \n", s3, s3)

	var u1, _ = strconv.ParseUint(s1, 10, 8)
	fmt.Println(u1)

	var i2, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Println(i2)

	// takes two parameters, base is implicit and not required.
	var f3, _ = strconv.ParseFloat(s3, 64)
	fmt.Println(f3)
}
