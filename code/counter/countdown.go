package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to time out for a Duration.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper implements Sleeper.
type DefaultSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

// Sleep sleeps for Duration.
func (ds DefaultSleeper) Sleep() {
	ds.sleep(ds.duration)
}

// Countdown counts down n times and finally prints an end message.
func Countdown(w io.Writer, times int, end string, s Sleeper) {
	for i := times; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, end)
}

func main() {
	sleeper := DefaultSleeper{2 * time.Second, time.Sleep}
	Countdown(os.Stdout, 4, "Boom!", sleeper)
}
