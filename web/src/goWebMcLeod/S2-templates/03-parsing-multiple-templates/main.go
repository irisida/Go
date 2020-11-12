package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("tpl-index.html", "tpl-about.html")

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-about.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
