package dictionary

// ErrNotFound is given when a key in a Dictionary is not found.
const ErrNotFound = OperationError("key not found")

// OperationError describes an Operationerror that happened while interacting with Dictionary.
type OperationError string

func (e OperationError) Error() string {
	return string(e)
}

// Dictionary holds keys that represents integers.
type Dictionary map[string]int

// Search returns a value for a specific key in Dictionary.
func (d Dictionary) Search(key string) (int, error) {
	val, ok := d[key]

	if !ok {
		return val, ErrNotFound
	}

	return val, nil
}

// Update updates the Dictionary with a (new) key-value entry.
func (d Dictionary) Update(key string, val int) {
	d[key] = val
}

// Delete removes a key-value pair from the Dictionary by its key.
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
