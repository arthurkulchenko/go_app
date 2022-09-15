package main

import(
	// "fmt"
	"net/http"
	"github.com/justinas/nosurf"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(response, request)
// 	})
// }

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie {
		HttpOnly: true,
		Path: "/",
		Secure: app.Env == "production",
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
