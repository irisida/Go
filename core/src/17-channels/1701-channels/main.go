package main

import "fmt"

func main() {

	var ch chan int

	// an uninitialised channel is nil
	fmt.Println(ch)

	// initliaise a channel
	ch = make(chan int)
	// recheck the output value
	fmt.Println(ch)

	// declare and initialise
	c := make(chan int)

	// the chanell operator is:  <-

	// send operation
	c <- 10

	// receive has the receiver, assignment, channel operator and channel
	num := <-c

	fmt.Println(<-c)
	fmt.Println(num)

	close(c)
}
