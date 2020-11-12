package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	/*
		use the glob operation to be able to point the program to
		a folder and dictate that all files ending with the .html
		extension are considered
	*/
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	beatles := map[string]string{
		"rhythm": "John",
		"bass":   "Paul",
		"lead":   "George",
		"drums":  "Ringo",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", beatles)
	if err != nil {
		log.Fatal(err)
	}

}
