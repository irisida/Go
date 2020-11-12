package main

import (
	"fmt"
)

func main() {
	src := "œ∑´®†¥¨ˆøπ" // 10 character/codepoint string
	rc := []rune(src)
	fmt.Printf("bytes: %v runes: %v", len(src), len(rc)) // 22 bytes and 10 runes

}
