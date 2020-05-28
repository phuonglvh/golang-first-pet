package models

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/phuonglvh/golang-first-pet/app/views"
)

// TemplateHandler modelling templates
type TemplateHandler struct {
	once     sync.Once
	Filename string
	template *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parse: ", t.Filename)
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("app/views", t.Filename)))
	})
	fmt.Println("Render: ", t.Filename)
	// t.template.Execute(w, r)
	err := views.ViewResolver.ExecuteTemplate(w, "chat", nil)
	fmt.Println((err))
}
