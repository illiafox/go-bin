package templates

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// // // //
const relative = "../../shared/templates"

// // // //
var (
	Error = newTemplate(relative + "/error/index.html")
	View  = newTemplate(relative + "/view/index.html")
)

// // // //

type Template struct {
	tmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data any) error {
	return t.tmpl.Execute(w, data)
}

func newTemplate(files ...string) *Template {
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("template: parse files '%s': %s", strings.Join(files, ","), err)
	}

	return &Template{tmpl: t}
}
