package controller

import (
	"html/template"
	"net/http"

	"github.com/phuonglvh/golang-first-pet/app/model"
)

// HomeHandler handle index page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := model.Page{
		Title:            "QR Code Generator",
		FormInputValue:   "",
		GeneratedCodeURI: "null",
	}
	t, _ := template.ParseFiles("template/generator.html")
	t.Execute(w, p)
}
