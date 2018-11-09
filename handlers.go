package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)


type comicManager interface {
	searchTranscripts(searchTerm string) []jsonComic
}

type jsonComic struct {
	Filename    string `json:"filename"`
	DisplayName string `json:"display_name"`
}

type searchResponse struct {
	CurrentPage int        `json:"current_page"`
	TotalPages int         `json:"total_pages"`
	Comics     []jsonComic `json:"comic"`

}

type searchRequest struct {
	Term       string `json:"term"`
	TargetPage int `json:"target_page"`
}

type searchHandler struct {
	storage comicManager
}

func (h searchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Interpret the search request

	var m searchRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//input search term
	if strings.TrimSpace(m.Term) == ""{
		w.WriteHeader(http.StatusOK)
		return
	}
	found := h.storage.searchTranscripts(m.Term)

	var result searchResponse
	result.Comics = found
	result.CurrentPage = 1
	result.TotalPages = 1

	// Response
	b, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
}

