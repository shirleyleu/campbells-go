package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type comicManagerTest struct {
	mock.Mock
}

func (m *comicManagerTest) searchTranscripts(searchTerm string) ([]dbComic, error) {
	args := m.Called(searchTerm)
	return args.Get(0).([]dbComic), args.Error(1)
}


func TestSearchHandler_ServeHTTP(t *testing.T) {
	b, _ := json.Marshal(searchRequest{})
	req := httptest.NewRequest("POST", "/json/search", bytes.NewBuffer(b))

	rr := httptest.NewRecorder()
	mockCM := new(comicManagerTest)
	handler := http.Handler(searchHandler{storage: mockCM})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockCM.AssertExpectations(t)
}

func TestSearchHandler_ServeHTTP_should_handle_empty_body(t *testing.T) {
	req := httptest.NewRequest("POST", "/json/search", nil)

	rr := httptest.NewRecorder()
	mockCM := new(comicManagerTest)
	handler := http.Handler(searchHandler{storage: mockCM})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockCM.AssertExpectations(t)

}

func TestSearchTranscript_empty_result(t *testing.T) {
	term := "gibberishthatwon'treturnanything"
	b, _ := json.Marshal(searchRequest{Term: term})
	req := httptest.NewRequest("POST", "/json/search", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()
	mockCM := new(comicManagerTest)
	mockCM.On("searchTranscripts",term).Return([]dbComic{}, nil)
	handler := http.Handler(searchHandler{storage: mockCM})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"current_page":1, "total_pages":1, "comic":[]}`, rr.Body.String())
	mockCM.AssertExpectations(t)
}

// Test for some results
// Test for results with more than 1 page
// Test when requested page is not 1
