package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /snippet/view/{id}", app.getHome)
	mux.HandleFunc("GET	/snippet/create", app.getHome)
	mux.HandleFunc("POST /snippet/create", app.getHome)

	return mux
}
