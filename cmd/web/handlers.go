package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Log the error.
		// app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		// Send internal server error to user.
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.URL.RequestURI()))
		// http.Error(w, "internal Server Error", http.StatusInternalServerError)
		app.serverError(w, r, err)
	}

}

func (app *application) getSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var message string = ("{\"snippet\":\"" + strconv.Itoa(id) + "\"}")
	//w.Write([]byte(message))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating a snippet"))
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creating new snippet..."))
}
