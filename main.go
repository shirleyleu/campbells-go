package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	cm, err := newComicManagerSQLite("./foxtrot.db")
	if err !=nil {
		log.Fatal(err)
	}
	r.Handle("/json/search", searchHandler{storage: cm})
	http.Handle("/", r)
	log.Print("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}