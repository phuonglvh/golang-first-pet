package model

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler modelling templates
type TemplateHandler struct {
	once     sync.Once
	Filename string
	template *template.Template
}

// ServeHTTP handles the HTTP request.
func (templateHandler *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parse: ", templateHandler.Filename)
	templateHandler.once.Do(func() {
		templateHandler.template = template.Must(template.ParseFiles(filepath.Join("template", templateHandler.Filename)))
	})
	fmt.Println("Render: ", templateHandler.Filename)
	templateHandler.template.Execute(w, r)
}
