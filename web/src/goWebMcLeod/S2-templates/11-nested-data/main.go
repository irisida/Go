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

type part struct {
	Manufacturer string
	Name         string
	Description  string
}

type model struct {
	Name  string
	Parts []part
}

type series struct {
	Models []model
}

func main() {

	p1 := part{
		Manufacturer: "TongFang",
		Name:         "PF14",
		Description:  "14 inch clamshell model",
	}
	p2 := part{
		Manufacturer: "Clevo",
		Name:         "CL15",
		Description:  "15 inch workstation clamshell",
	}
	p3 := part{
		Manufacturer: "DDR4",
		Name:         "DDR3200",
		Description:  "Fast RAM",
	}
	p4 := part{
		Manufacturer: "Nvidia",
		Name:         "3070",
		Description:  "bullion bar of GPU",
	}

	model1 := model{
		Name: "Ryzen 14",
	}
	model2 := model{
		Name: "Ryzen 7 15inch",
	}

	model1.Parts = append(model1.Parts, p1, p3, p4)
	model2.Parts = append(model2.Parts, p2, p3, p4)

	s2021 := series{}
	s2021.Models = append(s2021.Models, model1, model2)

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-index.html", s2021)
	if err != nil {
		log.Fatal(err)
	}

}
