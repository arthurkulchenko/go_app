package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.tmpl")
}

func About(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.tmpl")
}

func renderTemplate(response http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateName)
	error := parsedTemplate.Execute(response, nil)
	if error != nil {
		fmt.Println("Error parsing template:", error)
		return
	}
}
