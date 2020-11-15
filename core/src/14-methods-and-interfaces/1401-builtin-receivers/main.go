package main

import (
	"fmt"
	"time"
)

func main() {
	fp := fmt.Printf

	const day = 24 * time.Hour // creates a const of type time.Duration
	fp("day const type: %T \n", day)

	secs := day.Seconds() // this method is a receiver function
	fp("secs.type: %T secs.value:  %v\n", secs, secs)
}
