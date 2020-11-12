package main

import (
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
	bs := []byte("I am a slice of byte")

	bytesWritten, err := f.Write(bs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written: %d", bytesWritten)
}
