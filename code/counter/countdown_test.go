package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SleeperSpy struct {
	Calls []string
}

func (s *SleeperSpy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SleeperSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")

	return
}

type DurationSpy struct {
	durationSlept time.Duration
}

func (s *DurationSpy) Sleep(d time.Duration) {
	s.durationSlept = d
}

func TestCountdown(t *testing.T) {
	t.Run("counts down n times with end msg", func(t *testing.T) {
		spy := &SleeperSpy{}
		buffer := bytes.Buffer{}
		times := 3
		end := "Go!"

		Countdown(&buffer, times, end, spy)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("sleeps before every write", func(t *testing.T) {
		spy := &SleeperSpy{}

		Countdown(spy, 2, "Woohoo!", spy)

		got := spy.Calls
		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("sleeps for configured duration", func(t *testing.T) {
		spy := &DurationSpy{}
		duration := 2 * time.Second
		sleeper := DefaultSleeper{duration, spy.Sleep}

		sleeper.Sleep()

		got := spy.durationSlept
		want := duration

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
