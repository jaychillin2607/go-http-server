package main

import (
	"log"
	"net/http"

	server "github.com/jaychillin2607/go-http-server/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	return 123, true
}

func main() {
	myServer := &server.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", myServer))
}
