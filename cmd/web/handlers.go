package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"GoREST/internals/models"
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
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating a snippet"))
}

func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
