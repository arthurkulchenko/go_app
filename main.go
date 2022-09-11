package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const PORT_NUMBER = ":8080"

func home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.tmpl")
}

func about(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.tmpl")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}

func renderTemplate(response http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateName)
	error := parsedTemplate.Execute(response, nil)
	if error != nil {
		fmt.Println("Error parsing template:", error)
		return
	}
}
