package main

import (
	"fmt"
	"time"
)

// A buffered channel is one with a capacity set
// at the time where it is declared.

func main() {
	// make an unbuffered channel
	c := make(chan int, 3) // The sender of a buffered channel blocks only when there is
	// no empty slot on the channel.
	// The receive will block on the channel when its empty

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("inside go-routine -> buffered channel -> go-routine: %d\n", i)
			fmt.Println("\tsending data into channel")
			c <- i
			fmt.Println("\tafter sending data to channel")
		}
		close(c) // tells the receiving goroutines that all data has been sent.
	}(c)

	fmt.Println("inside main thread -> sleep for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("inside main thread -> start receiving data from channel")

	for v := range c {
		fmt.Printf("\t\t-> Main go-routine received: %d from the channel\n", v)
	}

}
