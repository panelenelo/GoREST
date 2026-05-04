package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port string = ":8181"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hilòu"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var message string = ("Viewing the snippet: " + strconv.Itoa(id))
	w.Write([]byte(message))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating a snippet"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("server running on localhost", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
