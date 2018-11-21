package racer

import (
	"net/http"
	"time"
)

// ErrTimeout describes a timeout Error.
const ErrTimeout = Error("request timed out")

// Error is a generic Error type to create more specific ones.
type Error string

func (e Error) Error() string {
	return string(e)
}

// Racer accepts 2 URLs and returns the fastest one to ping if timeout is not reached.
func Racer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
