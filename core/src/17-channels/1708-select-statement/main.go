package main

import (
	"fmt"
	"time"
)

func main() {
	// a select blocks until 1 of its cases can run
	// where one is found it then executes from a
	// random selection of cases which are ready.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "hello"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "World"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1: ", msg1)
		case msg2 := <-c2:
			fmt.Println("Recieved from c2: ", msg2)
		default:
			fmt.Println("no activity")
		}
	}
}
