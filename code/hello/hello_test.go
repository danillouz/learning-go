package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got '%s', want '%s'", got, want)
		}
	}

	t.Run("Greets in English with name", func(t *testing.T) {
		got := Hello("English", "Daniel")
		want := "Hello Daniel"

		assert(t, got, want)
	})

	t.Run("Default name to 'Go' in English", func(t *testing.T) {
		got := Hello("English", "")
		want := "Hello Go"

		assert(t, got, want)
	})

	t.Run("Greets in Dutch with name", func(t *testing.T) {
		got := Hello("Dutch", "Daniel")
		want := "Hallo Daniel"

		assert(t, got, want)
	})

	t.Run("Defaults name to 'Go' in Dutch", func(t *testing.T) {
		got := Hello("Dutch", "")
		want := "Hallo Go"

		assert(t, got, want)
	})
}
