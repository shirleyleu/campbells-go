package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type dbComic struct{
	date int
	transcript string
}

type comicManagerSQLite struct {
	c *sql.DB
}

func newComicManagerSQLite(source string) (*comicManagerSQLite, error){
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}
	return &comicManagerSQLite{db}, nil
}

func (m comicManagerSQLite) searchTranscripts(searchTerm string) ([]dbComic, error){

	return []dbComic{}, nil
}
