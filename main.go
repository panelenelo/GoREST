package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var token string
var randomGen *rand.Rand

func random(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer "+token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"value": randomGen.Intn(100)})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hiilóu")
		//io.WriteString(w, "hilöu\n")
	})
	fmt.Println("server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
