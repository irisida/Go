package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("check.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bs := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(f, bs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of bytes read: %d \n", numberBytesRead)
	log.Printf("Data read: %s \n", bs)

	f, err = os.Open("readall.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s \n", data)
	fmt.Printf("number of bytes read: %d \n", len(data))

	data, err = ioutil.ReadFile("readall.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s \n", data)
	//fmt.Printf("number of bytes read: %d \n", len(data))
}
