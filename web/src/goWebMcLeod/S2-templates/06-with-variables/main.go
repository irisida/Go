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
	var (
		err error
	)

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", "A value passed to a variable that is to be rendered")
	if err != nil {
		log.Fatal(err)
	}

}
