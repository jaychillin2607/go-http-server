package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/jaychillin2607/go-http-server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{}
	myServer := server.PlayerServer{Store: &store}
	player := "Pepper"

	myServer.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	myServer.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))
	myServer.ServeHTTP(httptest.NewRecorder(), server.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	myServer.ServeHTTP(response, server.NewGetScoreRequest(player))
	server.AssertStatus(t, response.Code, http.StatusOK)

	server.AssertResponseBody(t, response.Body.String(), "3")
}
