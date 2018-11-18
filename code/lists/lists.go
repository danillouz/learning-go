package main

// SumList return the sum of all elements in the provided list.
func SumList(list []int) int {
	var sum int

	for _, n := range list {
		sum += n
	}

	return sum
}

// SumListsInPlace returns a list that sums all elements in the provided lists.
func SumListsInPlace(lists ...[]int) []int {
	output := []int{}

	for _, list := range lists {
		sum := SumList(list)
		output = append(output, sum)
	}

	return output
}

func main() {
	//
}
