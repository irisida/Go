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

type beatle struct {
	Name string
	Role string
}

func main() {

	var beatles []beatle

	john := beatle{
		Name: "John Leonnon",
		Role: "rhythm",
	}
	paul := beatle{
		Name: "Paul McCartney",
		Role: "bass",
	}
	george := beatle{
		Name: "George Harrison",
		Role: "lead",
	}
	ringo := beatle{
		Name: "Ringo Starr",
		Role: "drums",
	}

	beatles = append(beatles, john, paul, george, ringo)

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", beatles)
	if err != nil {
		log.Fatal(err)
	}

}
