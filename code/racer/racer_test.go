package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}

	slowServer := httptest.NewServer(http.HandlerFunc(slowHandler))

	fastHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	fastServer := httptest.NewServer(http.HandlerFunc(fastHandler))

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	got := Racer(fastURL, slowURL)
	want := fastURL

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
