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
	if err != nil { log.Fatal("Cannot create template cache") }

	app.TemplateCache = templateCache
	app.PortNumber = PORT_NUMBER
	app.UseCache = false

	handlers.SetConfig(&app)
	fmt.Println(fmt.Sprintf("=======================\nStarting application on\nlocalhost%s\n=======================", app.PortNumber))
	server := &http.Server { Addr: app.PortNumber, Handler: Routes(&app) }
	err = server.ListenAndServe()
	log.Fatal(err)
}
