package main

import(
	"github.com/arthurkulchenko/go_app/pkg/config"
	"github.com/arthurkulchenko/go_app/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Routes(appP *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.RepositoryPointer.Home)
	mux.Get("/about", handlers.RepositoryPointer.About)
	return mux
}
