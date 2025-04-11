package main

import (
	"log"
	"net/http"
)

func main() {
	myServer := &PlayerServer{store: NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", myServer))
}
