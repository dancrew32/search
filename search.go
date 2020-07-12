package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", SearchHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
