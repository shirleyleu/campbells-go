package main

type dbComic struct{
	date int
	transcript string
}

type comicManagerSQLite struct {
}

func (m comicManagerSQLite) searchTranscripts(searchTerm string) ([]dbComic, error){

	return []dbComic{}, nil
}
