package main

import "fmt"

func main() {
	// declaring an array
	var nums [4]int

	// note as no values are specified each pof our int value elements will take the default zero value for int which is 0
	fmt.Printf("%v \n", nums)
	fmt.Printf("%#v \n", nums)

	var scores = [5]int{82, 88, 73, 69, 87}

	fmt.Printf("%v \n", scores)
	fmt.Printf("%#v \n", scores)

	// updating an element can be done easily by accessing the index of the element destined for change.
	scores[4] = 90
	scores[1] = 80

	fmt.Printf("%v \n", scores)
	fmt.Printf("%#v \n", scores)

	// using the ellipsis operator, saves defining the number of elements and calculates them at compile time.
	// they are still fixed but the operator offers syntactic sugar to the declaration process.
	elip := [...]int{90, 92, 93, 68, 71, 75, 83, 84, 81}

	fmt.Printf("%v \n", elip)
	fmt.Printf("%#v \n", elip)

	// using len to get the length
	fmt.Println(len(nums), len(scores), len(elip))
}
