package main

import (
	"fmt"
	"net/http"
	"github.com/arthurkulchenko/go_app/pkg/handlers"
	"github.com/arthurkulchenko/go_app/pkg/config"
	"log"
)

const PORT_NUMBER = ":8080"

func main() {
	var app config.AppConfig
	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = templateCache
	// app.UseCache = false
	app.UseCache = true

	// repo := handlers.NewRepo(&app)
	// handlers.NewRepo(&app)
	// handlers.NewHandlers(repo)
	handlers.SetConfig(&app)

	http.HandleFunc("/", handlers.RepositoryPointer.Home)
	http.HandleFunc("/about", handlers.RepositoryPointer.About)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}

// func NewRepo(appPointer *config.AppConfig) *Repository {
// 	RepositoryPointer = &Repository { AppPointer: appPointer, }
// }

// func NewHandlers(repositoryPointer *Repository) {
// 	RepositoryPointer = repositoryPointer
// }

