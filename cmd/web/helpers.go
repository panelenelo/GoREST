package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		// trace  = string(debug.Stack())
	)

	// Log the error.
	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri))
	// Send "internal server error" to user.
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not existe", page)
		app.serverError(w, r, err)
		return
	}

	buff := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buff, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buff.WriteTo(w)

}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}
