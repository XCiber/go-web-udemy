package render

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates from "templates" directory
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	if err := parsedTemplate.Execute(w, nil); err != nil {
		log.Printf("error parsing template: %v", err)
	}
}
