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

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatal(err)
	}
}
