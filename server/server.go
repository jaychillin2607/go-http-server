package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
}
type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, ok := p.Store.GetPlayerScore(player)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func GetPlayerScore(name string) (int, bool) {
	if name == "Pepper" {
		return 20, true
	}

	if name == "Floyd" {
		return 10, true
	}

	return 0, false
}
