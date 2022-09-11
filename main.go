package main

import (
	"fmt"
	"net/http"
	"errors"
)

const PORT_NUMBER = ":8080"

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "This is the home page")
}

func about(response http.ResponseWriter, request *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(response, fmt.Sprintf("This is the about page, and sum is %d", sum))
}

func divide(response http.ResponseWriter, request *http.Request) {
	division, divider := 10.0, 5.0
	result, error := divideValues(division, divider)
	if error != nil {
		fmt.Fprintf(response, "Cannot divide by zero")
		return
	}

	_, _ = fmt.Fprintf(response, fmt.Sprintf("%f divided by %f is %f", division, divider, result))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", divide)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}

func divideValues(x, y float64) (float64, error) {
	if y <= 0 {
		error := errors.New("Cannot divide by zero")
		return 0, error
	}

	result := x / y
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}
