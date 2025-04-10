package main

import (
	"log"
	"net/http"
)

func main() {
	myServer := &PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", myServer))
}
