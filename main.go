package main

import (
	"log"
	"net/http"
)

type InMemPlayerStore struct{}

func (i *InMemPlayerStore) GetPlayerScore(name string) int {
	if name == "Anna" {
		return 20
	}
	if name == "Dan" {
		return 10
	}

	return 0
}

func main() {
	// HandlerFunc allows the use of ordinary functions as HTTP handler
	// HandlerFunc has already implemented the ServeHTTP method and casts our
	// function with it meaning we have implemented the required "Handler"
	// for ListenAndServe
	//handler := http.HandlerFunc(PlayerServer)

	server := &PlayerServer{&InMemPlayerStore{}}
	// This is wrap calling an error and we will log it to the user, e.g.
	// if the port is already being listened on
	log.Fatal(http.ListenAndServe(":5000", server))
}
