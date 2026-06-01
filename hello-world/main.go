package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":3000"

var tmplCache = template.Must(template.ParseGlob("./templates/*.tmpl"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server is starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
