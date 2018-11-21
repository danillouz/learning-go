package racer

import (
	"net/http"
)

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}

// Racer accepts 2 URLs and returns the fastest one to ping.
func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
