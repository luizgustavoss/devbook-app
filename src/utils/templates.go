package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates loads html pages to template variable
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// RenderTemplate renders a specific html template
func RenderTemplate(w http.ResponseWriter, template string, data interface{}){
	templates.ExecuteTemplate(w, template, data)
}