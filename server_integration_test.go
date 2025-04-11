package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	myServer := NewPlayerServer(store)
	player := "Pepper"

	myServer.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	myServer.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	myServer.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	response := httptest.NewRecorder()
	myServer.ServeHTTP(response, NewGetScoreRequest(player))
	AssertStatus(t, response.Code, http.StatusOK)

	AssertResponseBody(t, response.Body.String(), "3")
}
