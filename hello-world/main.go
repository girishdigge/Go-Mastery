package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 5)
	fmt.Fprintf(w, "This is the about page and 2+5: %d", sum)
}

func Devide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(10.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 10.0, 0.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Devide)

	fmt.Printf("Server is starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
