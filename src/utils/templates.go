package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadTemplate add the HTML template on var template
func LoadTemplate() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate render a html page on screen
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
