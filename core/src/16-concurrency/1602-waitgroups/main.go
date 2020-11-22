package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")

	for i := 0; i < 5; i++ {
		fmt.Printf("\tloop [i] is: %v\n", i)
	}
	time.Sleep(time.Second * 2) // simulate an expensive task

	fmt.Println("f1 exiting")
	wg.Done()
	// may be called as (*wg).Done() also
}

func f2() {
	fmt.Println("f2 started")

	for i := 5; i < 10; i++ {
		fmt.Printf("\tloop [i] is: %v\n", i)
	}

	fmt.Println("f2 exiting")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1) // number of goRoutines to wait for

	go f1(&wg)
	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	f2()

	wg.Wait()
	fmt.Println("main execution concluded")
}
