package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Log the error.
		log.Print(err.Error())
		// Send internal server error to user.
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal Server Error", http.StatusInternalServerError)
	}

}

func getSnippetView(w http.ResponseWriter, r *http.Request) {
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

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating a snippet"))
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creating new snippet..."))
}
