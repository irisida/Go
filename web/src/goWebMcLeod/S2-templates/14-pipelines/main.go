package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqrt(x float64) float64 {
	return math.Sqrt(float64(x))
}

var tpl *template.Template
var fm = template.FuncMap{
	"fDbl":  double,
	"fSq":   square,
	"fSqrt": sqrt,
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", 3)
	if err != nil {
		log.Fatal(err)
	}

}
