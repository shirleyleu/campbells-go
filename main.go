package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/json/search", searchHandler{storage: comicManagerSQLite{}})
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}