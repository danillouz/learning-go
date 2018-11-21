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
	t.Run("pings URLs and returns the fastest to resolve", func(t *testing.T) {
		slowServer := newDelayedServer(10 * time.Millisecond)
		defer slowServer.Close()

		fastServer := newDelayedServer(0)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(fastURL, slowURL, 20*time.Millisecond)
		want := fastURL

		if err != nil {
			t.Fatalf("got error %v, want URL", err)
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("return error when ping takes longer than timeout", func(t *testing.T) {
		slowServer := newDelayedServer(10 * time.Millisecond)
		defer slowServer.Close()

		fastServer := newDelayedServer(10 * time.Millisecond)
		defer slowServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		timeout := 5 * time.Millisecond
		_, err := Racer(slowURL, fastURL, timeout)

		if err != ErrTimeout {
			t.Errorf("got %v, want error %v", err, ErrTimeout)
		}
	})
}
