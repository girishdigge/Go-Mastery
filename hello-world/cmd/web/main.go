package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/girishdigge/hello-world/pkg/config"
	"github.com/girishdigge/hello-world/pkg/handlers"
	"github.com/girishdigge/hello-world/pkg/render"
)

const portNumber = ":3000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Server is starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
