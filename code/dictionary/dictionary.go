package dictionary

import (
	"errors"
)

// ErrorNotFound is given when a key in a Dictionary is not found
var ErrorNotFound = errors.New("key not found")

// Dictionary holds keys that represents integers.
type Dictionary map[string]int

// Search returns a value for a specific key in Dictionary.
func (d Dictionary) Search(key string) (int, error) {
	val, ok := d[key]

	if !ok {
		return val, ErrorNotFound
	}

	return val, nil
}
