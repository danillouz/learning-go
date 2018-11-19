package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	key := "two"
	got := Search(dict, key)
	want := dict[key]

	if got != want {
		t.Errorf("got %v, want %v, for dict %#+v and key %s", got, want, dict, key)
	}
}
