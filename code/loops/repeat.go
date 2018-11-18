package loops

// RepeatFive accepts an input string and repeats it 5 times.
func RepeatFive(in string) string {
	var out string

	for i := 0; i < 5; i++ {
		out += in
	}

	return out
}
