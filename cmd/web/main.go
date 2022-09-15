package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/arthurkulchenko/go_app/pkg/config"
	"github.com/arthurkulchenko/go_app/pkg/handlers"
	"log"
	"net/http"
	"time"
)

const PORT_NUMBER = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.Env = "development"
	sessionInitiated := scs.New()
	sessionInitiated.Lifetime = 24 * time.Hour
	sessionInitiated.Cookie.Persist = true
	sessionInitiated.Cookie.SameSite = http.SameSiteLaxMode
	sessionInitiated.Cookie.Secure = app.Env == "production"
	app.Session = sessionInitiated

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
