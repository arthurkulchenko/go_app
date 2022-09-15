package handlers

import (
	// "fmt"
	"net/http"
	"html/template"
	"log"
	"path/filepath"
	"bytes"
	"github.com/arthurkulchenko/go_app/pkg/config"
	// "github.com/arthurkulchenko/go_app/pkg/models"
)

func (m *Repository) Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home.page.tmpl", &TemplateData{})
}

func (m *Repository) About(response http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"
	renderTemplate(response, "about.page.tmpl", &TemplateData { StringMap: stringMap })
}

var RepositoryPointer *Repository
var appConfigP *config.AppConfig

type Repository struct {
	AppConfigPointer *config.AppConfig
}

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Watrning string
	Error string
}

func SetConfig(appConfigPointer *config.AppConfig) {
	// RepositoryPointer = &Repository { AppConfigPointer: appConfigPointer, }
	appConfigP = appConfigPointer
}

func addDefaultData(templateDataPointer *TemplateData) *TemplateData {
	return templateDataPointer
}

func renderTemplate(response http.ResponseWriter, templateName string, templateData *TemplateData) {
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
	templateData = addDefaultData(templateData)
	err := cachedTemplate.Execute(buffer, templateData)
	if err != nil { log.Println(err) }

	// render the template
	_, err = buffer.WriteTo(response)
	if err != nil { log.Println(err) }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all files with *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil { return myCache, err }
	// range trough all files with template extention
	for _, page := range pages {
		name := filepath.Base(page) // returs last element from '/'
		parsedTemplatePointer, err := template.New(name).ParseFiles(page)
		if err != nil { return myCache, err }

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil { return myCache, err }
		if len(layouts) > 0 {
			parsedTemplatePointer, err = parsedTemplatePointer.ParseGlob("./templates/*.layout.tmpl")
			if err != nil { return myCache, err }
		}
		myCache[name] = parsedTemplatePointer
	}
	return myCache, nil
}
