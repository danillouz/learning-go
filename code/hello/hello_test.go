package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got '%s', expected '%s'", got, want)
		}
	}

	t.Run("Appends name", func(t *testing.T) {
		got := Hello("Daniel")
		want := "Hello Daniel"

		assert(t, got, want)
	})

	t.Run("Default name to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assert(t, got, want)
	})
}
