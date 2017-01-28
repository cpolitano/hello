package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template // init tpl as pointer to template

func init() {
	// tpl is a container holding all the templates
	tpl = template.Must(template.ParseGlob("views/*"))
	// ParseGlob returns pointer to template and err
	// Must takes pointer to template and err and does
	// error checking for us
}

func main() {
	// ParseFiles parses one or more named files
	// tpl, err := template.ParseFiles("template.gohtml")

	err := tpl.Execute(os.Stdout, nil) // pipes template to new file
	if err != nil {
		log.Fatalln(err)
	} 

	// execute a specific template
	// args are destination, template name, and data
	err = tpl.ExecuteTemplate(os.Stdout, "greet.gohtml", nil) 
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "updates.gohtml", nil) 
	if err != nil {
		log.Fatalln(err)
	}

	// nf, err := os.Create("index.html") // create a new file
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer nf.Close()

}


