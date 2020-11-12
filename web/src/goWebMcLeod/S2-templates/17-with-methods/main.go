package main

import (
	"html/template"
	"log"
	"os"
)

type composer struct {
	person
	IsPublished bool
}

type person struct {
	Name string
	Role string
}

func (p person) SomeMethod() string {
	return "SomeMethod was run"
}

func (p person) SomeOtherMethod() string {
	return "SomeOtherMethod was run"
}

var tpl *template.Template

func init() {
	/*
		passes template.New() to must as the static compilaion demands that the func
		be there at compile time. Creating a standard template and attaching the
		funcs afterwards will not work.
	*/
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	jl := composer{
		person{
			Name: "John Lennon",
			Role: "Rhyhtm guitar",
		},
		false,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", jl)
	if err != nil {
		log.Fatal(err)
	}

}
