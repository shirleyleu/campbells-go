package main

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type dbComic struct{
	date int
}

type comicManagerSQLite struct {
	db *sql.DB
}

func newComicManagerSQLite(source string) (*comicManagerSQLite, error){
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, err
	}
	return &comicManagerSQLite{db}, nil
}

func (m *comicManagerSQLite) searchTranscripts(searchTerm string) ([]dbComic, error){
	query := `SELECT date FROM search_transcripts WHERE transcript MATCH $1`
	rows, err := m.db.QueryContext(context.TODO(), query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var found []dbComic
	for rows.Next(){
		var date int
		err = rows.Scan(&date)
		if err != nil {
			return nil, err
		}
		found = append(found, dbComic{date})
	}
	return found, nil
}
