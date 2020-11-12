package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	a, b := 'a', 'b'

	fmt.Printf("%T %v \n", a, a) // shows the type int32 which is the type or alias type fpr a rune
	fmt.Printf("%T %v \n", b, b)

	str := "Keeƒ" // note that character at str[3] takes up more bytes.
	fmt.Println(len(str))
	fmt.Println("Byte (not rune) at position 2", str[2])

	// fixed increment has unexpected results for some runes
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n", strings.Repeat("-", 20))

	// increment by rune size
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n", strings.Repeat("-", 20))

	// standard character pattern
	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
