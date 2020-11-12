package main

import (
	"fmt"
	"strings"
)

func main() {

	pt := fmt.Println
	pf := fmt.Printf

	res := strings.Contains("Source value to be searched", "value") // source, searchterm
	pt(res)

	// contains
	res = strings.ContainsAny("Beatles", "John") // returns false as no matching letters
	pt(res)
	res = strings.ContainsAny("Beatles", "George") // matches on the `e` therefore returns true
	pt(res)

	// runes
	res = strings.ContainsRune("Golang", 'g')
	pt(res)

	// appearances
	count := strings.Count("Mississippi", "s")
	pt(count)

	// casing
	pt(strings.ToLower("GOLANG iS A GoOD LAnGuAGe TO wOrk WiTH"))
	pt(strings.ToUpper("GOLANG iS A GoOD LAnGuAGe TO wOrk WiTH"))

	// repeat
	pt(strings.Repeat("-", 5))

	// replace
	str := strings.Replace("192.168.1.2", ".", ":", -1) // negative value here hotrods string to do all replacements
	pt(str)

	// split
	sp := strings.Split("One two three four five once I caught a fish alive", " ")
	pf("%T \n", sp)
	pf("%#v \n", sp)

	s := []string{"192", "168", "1", "1"}
	js := strings.Join(s, ".")
	pt(js)

	// strings.TrimSpace(sourceWithLeadingSpaces)
	// strings.Trim(source, trimlist)
}
