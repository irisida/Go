package main

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func main() {
	// create two books
	book1 := book{title: "The Devine Comedy", author: "Dante Aligheri", year: 1320}
	book2 := book{title: "MacBeth", author: "William Shakespear", year: 1606}

	// print representation usual the struct we created instead of multiple variables
	fmt.Println(book1)
	fmt.Println(book2)

	// comparison
	fmt.Println(book1 == book2)

	book3 := book{title: "MacBeth", author: "William Shakespear", year: 1606}
	fmt.Println(book2 == book3)

	// where the fields are matched the struct comparison will return true.

	// copies
	myBook := book3
	myBook.year = 2020
	fmt.Println(book3)
	fmt.Println(myBook)

	// note that copies are not memory efficient in the sense of
	// sharing the same same underlying memory, they are independent
	// copies of the data and a change in one os not reflected in
	// the other.
}
