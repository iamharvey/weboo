package internal

import (
	"html/template"
	"net/http"
)

// renderTemplate renders html template files. Handler function should use it to render template file content.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HomeHandler renders home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}
