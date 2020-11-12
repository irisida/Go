package main

import "fmt"

func main() {

	// copy works with length in mind.
	// where length is equalled with the source, a deep copy os created
	src := []int{10, 20, 30, 40, 50}
	dst := make([]int, len(src))
	res := copy(dst, src)
	fmt.Println(src, dst, res)

	// where length of destination differs, only the matching elements will be copied.
	// in this instance we take the first 3.
	nDst := make([]int, 3)
	res = copy(nDst, src)
	fmt.Println(src, nDst, res)

	// if a copy destination has a length of zero the copy operation will not error
	// it will copy nothing to the new structure
	zDst := make([]int, 0)
	res = copy(zDst, src)
	fmt.Println(src, zDst, res)
}
