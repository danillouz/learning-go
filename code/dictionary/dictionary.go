package dictionary

// Search returns a value for a specific key in a dictionary.
func Search(dict map[string]int, key string) int {
	return dict[key]
}
