package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	want := "hello"

	Greet(&buffer, want)

	got := buffer.String()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
