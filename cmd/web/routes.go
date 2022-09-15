package main

import(
	"net/http"
	"github.com/arthurkulchenko/go_app/pkg/config"
	"github.com/bmizerany/pat"
	"github.com/arthurkulchenko/go_app/pkg/handlers"
)

func Routes(appP *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.RepositoryPointer.Home))
	mux.Get("/about", http.HandlerFunc(handlers.RepositoryPointer.About))
	return mux
}
