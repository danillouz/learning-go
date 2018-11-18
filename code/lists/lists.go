package main

// SumList return the sum of all elements in the provided list.
func SumList(list []int) int {
	var sum int

	for _, n := range list {
		sum += n
	}

	return sum
}

func main() {
	//
}
