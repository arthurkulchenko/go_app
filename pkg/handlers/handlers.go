package handlers

import (
	// "fmt"
	"net/http"
	"html/template"
	"log"
	"path/filepath"
	"bytes"
	"github.com/arthurkulchenko/go_app/pkg/config"
)

func (m *Repository) Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.tmpl")
}

func (m *Repository) About(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about.page.tmpl")
}

var RepositoryPointer *Repository
var appConfigP *config.AppConfig

type Repository struct {
	AppConfigPointer *config.AppConfig
}

// NewRepo creates a new repository
// func NewRepo(appConfigPointer *config.AppConfig) *Repository {
// func NewRepo(appConfigPointer *config.AppConfig) {
// 	RepositoryPointer = &Repository { AppConfigPointer: appConfigPointer, }
// 	// return &Repository {
// 	// 	AppConfigPointer: appConfigPointer,
// 	// }
// }

// New sets the repository for handlers
// func NewHandlers(repositoryPointer *Repository) {
// 	RepositoryPointer = repositoryPointer
// }

func SetConfig(appConfigPointer *config.AppConfig) {
	// RepositoryPointer = &Repository { AppConfigPointer: appConfigPointer, }
	appConfigP = appConfigPointer
}

func renderTemplate(response http.ResponseWriter, templateName string) {
	var templateCache map[string]*template.Template
	if appConfigP.UseCache {
		templateCache = appConfigP.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	// get requested template
	cachedTemplate, exists := templateCache[templateName]
	if !exists {
		log.Fatal("Could not get template cache")
	}

	buffer := new(bytes.Buffer)
	err := cachedTemplate.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buffer.WriteTo(response)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all files with *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range trough all files with template extention
	for _, page := range pages {
		name := filepath.Base(page) // returs last element from '/'
		parsedTemplatePointer, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			parsedTemplatePointer, err = parsedTemplatePointer.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = parsedTemplatePointer
	}
	return myCache, nil
}

// var templateCache = make(map[string]*template.Template)

// func renderTemplate(response http.ResponseWriter, templateName string) {
// 	var err error
// 	templates := []string { fmt.Sprintf("./templates/%s", templateName), "./templates/base.layout.tmpl", }
// 	// check to see if we already hae the template in the cache
// 	_, inMap := templateCache[templateName]
// 	if inMap {
// 		log.Println("using cached template")
// 	} else {
// 		log.Println("creating template")
// 		tmpl, err := template.ParseFiles(templates...)

// 		if err != nil {
// 			log.Println(err)
// 		} else {
// 			templateCache[templateName] = tmpl
// 		}
// 	}

// 	err = templateCache[templateName].Execute(response, nil)
// 	if err != nil { log.Println(err) }
// }
