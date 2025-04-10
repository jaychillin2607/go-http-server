package main

import (
	"log"
	"net/http"
)

func main() {
	myServer := &PlayerServer{Store: NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", myServer))
}
