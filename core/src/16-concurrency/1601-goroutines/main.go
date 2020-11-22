package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Main execution started...")
	fmt.Println("No of CPUs : ", runtime.NumCPU())
	fmt.Println("No of goroutines: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)

	fmt.Println("Max Procs", runtime.GOMAXPROCS(0))

}
