package main

import (
	"log"
	"os"
	"text/template"
	"strings"
)

var tpl *template.Template // init tpl as pointer to template

type friend struct {
	Name string
	Greeting string
}

// FuncMap registers functions to pass into template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree, 
}

func init() {
	// tpl is a container holding all the templates
	// tpl = template.Must(template.ParseGlob("views/*"))
	// ParseGlob returns pointer to template and err
	// Must takes pointer to template and err and does
	// error checking for us

	// New creates pointer to template
	// add Funcs before Parse
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("views/*"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	// ParseFiles parses one or more named files
	// tpl, err := template.ParseFiles("template.gohtml")

	// err := tpl.Execute(os.Stdout, nil) // pipes template to new file
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// execute a specific template
	// args are destination, template name, and data
	// value of type STRING
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", "Clare")
	if err != nil {
		log.Fatalln(err)
	}

	// value of type MAP
	greetings := map[string]string{
		"English": "Hello",
		"Hawaiian": "Aloha",
		"Italian": "Ciao",
		"Arabic": "Marhaba",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "greet.gohtml", greetings)
	if err != nil {
		log.Fatalln(err)
	}

	// value of type SLICE
	updates := []string{"news 1", "news 2", "news 3", "news 4"}
	err = tpl.ExecuteTemplate(os.Stdout, "updates.gohtml", updates)
	if err != nil {
		log.Fatalln(err)
	}

	// value of type STRUCT
	anisa := friend{
		"Anisa",
		"Hey",
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", anisa)
	if err != nil {
		log.Fatalln(err)
	}

	// value of type SLICE of STRUCTS
	david := friend{
		"David",
		"Hi",
	}
	zoe := friend{
		"Zoe",
		"yo",
	}
	friends := []friend{anisa, david, zoe}
	err = tpl.ExecuteTemplate(os.Stdout, "friends.gohtml", friends)
	if err != nil {
		log.Fatalln(err)
	}

	// same but with functions
	err = tpl.ExecuteTemplate(os.Stdout, "func.gohtml", friends)
	if err != nil {
		log.Fatalln(err)
	}

	// nf, err := os.Create("index.html") // create a new file
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer nf.Close()

}
