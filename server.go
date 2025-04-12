package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() League
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.HandleFunc("/league", p.leagueHandler)
	router.HandleFunc("/players/", p.playersHandler)

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := p.store.GetLeague()
	w.Header().Set("Content-Type", jsonContentType)

	err := json.NewEncoder(w).Encode(leagueTable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "internal server error"}`)
		return
	}
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score, ok := p.store.GetPlayerScore(player)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
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
