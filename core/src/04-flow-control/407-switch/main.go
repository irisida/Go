package main

import (
	"fmt"
	"time"
)

func main() {

	hour := time.Now().Hour()
	min := time.Now().Minute()

	fmt.Printf("Current time is %v:%v, ", hour, min)

	// switch controller
	switch {
	case hour < 12:
		fmt.Println("that makes it Morning time.")
	case hour < 18:
		fmt.Println("it's the Afternoon.")
	default:
		fmt.Println("good Evening.")
	}
}
