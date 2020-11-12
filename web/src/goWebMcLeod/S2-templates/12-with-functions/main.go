package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"lc": strings.ToLower,
	"f3": firstThree,
}

func init() {
	/*
		passes template.New() to must as the static compilaion demands
		that the func be there at compile time. Creating a standard template
		and attaching the funcs afterwards will not work.
	*/
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type thing struct {
	Name string
	Desc string
}

type data struct {
	Things []thing
}

func main() {

	t1 := thing{
		Name: "One",
		Desc: "Singular",
	}
	t2 := thing{
		Name: "Two",
		Desc: "Has twice the value of one",
	}
	t3 := thing{
		Name: "Three",
		Desc: "Its the magic number",
	}
	t4 := thing{
		Name: "Four",
		Desc: "Rhymes with floor",
	}

	data := data{}
	data.Things = append(data.Things, t1, t2, t3, t4)

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", data)
	if err != nil {
		log.Fatal(err)
	}

}
