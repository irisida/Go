package main

import (
	"fmt"
	"time"
)

func main() {
	// make an unbuffered channel
	c := make(chan int)

	go func(c chan int) {
		fmt.Println("inside go-routine -> sending data into channel")
		c <- 100
		fmt.Println("inside go-routine -> after sending data to channel")
	}(c)

	fmt.Println("inside main thread -> sleep for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("inside main thread -> start receiving data from channel")
	d := <-c
	fmt.Println("inside main thread -> received data: ", d)

}
