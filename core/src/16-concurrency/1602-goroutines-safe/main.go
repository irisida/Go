package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 started")

	for i := 0; i < 5; i++ {
		fmt.Printf("\tloop [i] is: %v\n", i)
	}

	fmt.Println("f1 exiting")
}

func f2() {
	fmt.Println("f2 started")

	for i := 5; i < 10; i++ {
		fmt.Printf("\tloop [i] is: %v\n", i)
	}

	fmt.Println("f2 exiting")
}

func main() {

	fmt.Println("Main execution started...")
	fmt.Println("No of CPUs : ", runtime.NumCPU())
	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)

	fmt.Println("Max Procs", runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("main execution concluded")
}
