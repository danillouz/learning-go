package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newDelayedServer(delay time.Duration) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	return server
}

func TestRacer(t *testing.T) {
	slowServer := newDelayedServer(10 * time.Millisecond)
	defer slowServer.Close()

	fastServer := newDelayedServer(0)
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	got := Racer(fastURL, slowURL)
	want := fastURL

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
