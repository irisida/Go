package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("tpl-main.html")

	if err != nil {
		log.Fatal(err)
	}

	// create a new file, or log err if thrown
	// defer close to the end of the function
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error - creating file: ", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
