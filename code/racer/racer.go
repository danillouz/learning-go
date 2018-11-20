package racer

import (
	"net/http"
	"time"
)

// Racer accepts 2 URLs and returns the fastest one to resolve.
func Racer(a, b string) string {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)

	if durationA < durationB {
		return a
	}

	return b
}
