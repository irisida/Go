package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//pf := fmt.Printf
	p := fmt.Println

	// create a pointer to a new file
	var nf *os.File
	var err error

	nf, err = os.Create("./a.txt")
	if err != nil {
		log.Fatal(err) // threadsafe and adds timings. The more idiomatic err handler
	}

	// truncate a file. Truncate() takes a (path)filename and a size, if zero it truncates all content
	err = os.Truncate("a.txt", 0)

	// close a file
	nf.Close()

	// open a file
	file, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// the fileInfo object
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p("Filename: ", fileInfo.Name())
	p("fileSize: ", fileInfo.Size())
	p("Modified: ", fileInfo.ModTime())
	p("Is a dir: ", fileInfo.IsDir())
	p("perms   : ", fileInfo.Mode())

	// fileinfo also returns an err, so in the case of a non existent file we have an
	// error we can work with and process.
	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}

	op := "a.txt"
	np := "c.txt"
	err = os.Rename(op, np)
	if err != nil {
		log.Fatal(err)
	}

	// remove a file
	err = os.Remove("c.txt")
	if err != nil {
		log.Fatal(err)
	}
}
