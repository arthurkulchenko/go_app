package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

var templateCache = make(map[string]*template.Template)

func Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.tmpl")
}

func About(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.tmpl")
}

func renderTemplate(response http.ResponseWriter, templateName string) {
	var err error
	templates := []string { fmt.Sprintf("./templates/%s", templateName), "./templates/base_layout.tmpl", }
	// check to see if we already hae the template in the cache
	_, inMap := templateCache[templateName]
	if inMap {
		log.Println("using cached template")
	} else {
		log.Println("creating template")
		tmpl, err := template.ParseFiles(templates...)

		if err != nil {
			log.Println(err)
		} else {
			templateCache[templateName] = tmpl
		}
	}

	err = templateCache[templateName].Execute(response, nil)
	if err != nil { log.Println(err) }
}
