package racer

import (
	"net/http"
	"time"
)

func getRequestDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)

	return duration
}

// Racer accepts 2 URLs and returns the fastest one to resolve.
func Racer(a, b string) string {
	durationA := getRequestDuration(a)
	durationB := getRequestDuration(b)

	if durationA < durationB {
		return a
	}

	return b
}
