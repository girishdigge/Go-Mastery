package render

import (
	"log"
	"net/http"
	"text/template"
)

var tmplCache = template.Must(template.ParseGlob("./templates/*.tmpl"))

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	err := tmplCache.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		log.Printf("Template execution error for %s: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
