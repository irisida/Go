package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var tpl *template.Template
var fm = template.FuncMap{
	"fDateMDY": monthDayYear,
}

func init() {
	/*
		passes template.New() to must as the static compilaion demands that the func
		be there at compile time. Creating a standard template and attaching the
		funcs afterwards will not work.
	*/
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", time.Now())
	if err != nil {
		log.Fatal(err)
	}

}
