package main

import "fmt"

func main() {
	s1 := "Some text value" // a string literal
	fmt.Println(s1)

	// where a double quote is required insode a string it must be escaped
	s2 := "Some quoted text \"notable quote!\""
	fmt.Println(s2)

	// back ticks are also allowed. These are what is called a raw string,
	// escaped sequences will have no effect here and show as their text
	// representation instead.
	s3 := `This is another example of "some quoted text"`
	fmt.Println(s3)

	// a string is in fact a slice of bytes so accessing the index of
	// a string character will not return the character but the utf-8
	// value of it instead.
	s4 := "Utf-8 encoded byte slice"
	val := s4[6]     // expects the character 'e'
	fmt.Println(val) // prints 101, unicode of `e` is 101

}
