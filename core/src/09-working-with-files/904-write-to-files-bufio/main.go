package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("b.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufferedWriter := bufio.NewWriter(f)

	// create a byteslice of data to be written.
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (Not file) : %d \n", bytesWritten)

	// buffer default size is 4096. We don't want to write to disk with every byte as that is very
	// inefficient. So we create a buffer that can optimise the writes.
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d \n", bytesAvailable)

	// create a new string of text data
	bytesWritten, err = bufferedWriter.WriteString("Some string of text")
	if err != nil {
		log.Fatal(err)
	}

	// the .Buffered() tell us how much of the 4096 is allocated for a write.
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("bytes buffered: %d \n", unflushedBufferSize)

	// writes the file out.
	bufferedWriter.Flush()
}
