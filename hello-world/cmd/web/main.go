package main

import (
	"fmt"
	"net/http"

	"github.com/girishdigge/hello-world/pkg/handlers"
)

const portNumber = ":3000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server is starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
