package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("b.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write operations
	bs := []byte("Programming with Go is great")

	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
