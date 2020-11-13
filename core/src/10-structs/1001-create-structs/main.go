package main

import "fmt"

type book struct {
	title  string
	author string
	year   int
}

func main() {
	// collection of unrelated variables example
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "MacBeth", "William Shakespear", 1606

	fmt.Println("book1: ", title1, author1, year1)
	fmt.Println("book2: ", title2, author2, year2)

	// examples of structured data
	book1 := book{title: "The Devine Comedy", author: "Dante Aligheri", year: 1320}
	book2 := book{title: "MacBeth", author: "William Shakespear", year: 1606}

	// print representation usual the struct we created instead of multiple variables
	fmt.Println(book1)
	fmt.Println(book2)
}
