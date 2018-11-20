package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	assertError := func(t *testing.T, got, want error) {
		if got == nil {
			t.Fatalf("got %v, want err", got)
		}

		if got != want {
			t.Errorf("got err %v, want err %v", got, want)
		}
	}

	t.Run("match", func(t *testing.T) {
		key := "two"
		got, _ := dict.Search(key)
		want := dict[key]

		if got != want {
			t.Errorf("got %v, want %v, for dict %#+v and key %s", got, want, dict, key)
		}
	})

	t.Run("no match", func(t *testing.T) {
		_, err := dict.Search("nokey")

		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{}
	key := "one"
	want := 1
	dict.Update(key, want)

	got, err := dict.Search(key)

	if err != nil {
		t.Fatalf("got err %v, want value %v", err, want)
	}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	key := "one"
	dict := Dictionary{}

	dict[key] = 1
	dict.Delete(key)
	_, err := dict.Search(key)

	if err != ErrNotFound {
		t.Fatalf("got %v, want error %v", err, ErrNotFound)
	}
}
