package main

import (
	"fmt"
	"net/http"
)

const PORT_NUMBER = ":8080"

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "This is the home page")
}

func about(response http.ResponseWriter, request *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(response, fmt.Sprintf("This is the about page, and sum is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
