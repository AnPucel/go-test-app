package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Stub player store
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Anna": 20,
			"Dan":  10,
		},
	}

	// Create an instance of player server
	server := &PlayerServer{&store}

	t.Run("returns Anna's score", func(t *testing.T) {
		request := newGetScoreRequest("Anna")

		// net/http/httptest has a spy already made to spy on the responses
		response := httptest.NewRecorder()

		//PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Dan's score", func(t *testing.T) {
		request := newGetScoreRequest("Dan")

		// net/http/httptest has a spy already made to spy on the responses
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 for unknown player", func(t *testing.T) {
		request := newGetScoreRequest("Randal")

		// net/http/httptest has a spy already made to spy on the responses
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}

	// Create an instance of player server
	server := &PlayerServer{&store}

	t.Run("returns Anna's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "players/Anna", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

/*
func newStoreScoreRequest(name string, numWins int) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("players/%s", name, nil))
	return req
}
*/

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
