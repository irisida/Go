package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("lines.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// read the first line example
	// success := scanner.Scan()
	// if success == false {
	// 	err = scanner.Err()
	// 	if err == nil {
	// 		log.Println("reached EOF")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Println("first line found: ", scanner.Bytes())

	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
