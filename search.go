package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type SearchResponse struct {
	Data string `json:"data"`
}

func SearchHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	search, ok := query["search"]
	searchResponse := &SearchResponse{}
	if ok && len(search) == 1 {
		toSearch := search[0]
		// TODO(DAN): hit redisearch,
		// searchResults := ...
		searchResponse.Data = toSearch
	}
	bytes, _ := json.Marshal(searchResponse)
	writer.Write(bytes)
}

type IndexHandler struct{}

func (handler IndexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./index.html")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/search", SearchHandler)
	router.PathPrefix("/").Handler(IndexHandler{})
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 50 * time.Millisecond,
		ReadTimeout:  50 * time.Millisecond,
	}
	log.Fatal(server.ListenAndServe())
}
