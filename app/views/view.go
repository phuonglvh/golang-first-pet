package views

import (
	"fmt"
	"html/template"
)

// ViewResolver will handle template calls
var ViewResolver *template.Template

func init() {
	templ, err := template.ParseGlob("app/views" + "/**/*.gohtml")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Defined templates: %s", templ.DefinedTemplates())
		ViewResolver = templ
	}
}
