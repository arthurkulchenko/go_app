package main

import (
	"fmt"
	"net/http"
	"github.com/arthurkulchenko/go_app/pkg/handlers"
	// "html/template"
)

const PORT_NUMBER = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
