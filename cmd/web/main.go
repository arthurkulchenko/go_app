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

	handlers.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", PORT_NUMBER))
	_ = http.ListenAndServe(PORT_NUMBER, nil)
}
