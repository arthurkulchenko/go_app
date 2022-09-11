package main

import (
	"fmt"
	"net/http"
	// "html/template"
)

const PORT_NUMBER = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
