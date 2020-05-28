package views

import (
	"html/template"
	"log"
)

// ViewResolver will handle template calls
var ViewResolver *template.Template

func init() {
	templ, err := template.ParseGlob("app/views" + "/**/*.gohtml")
	if err != nil {
		panic(err)
	} else {
		log.Println(templ.DefinedTemplates())
		ViewResolver = templ
	}
}
