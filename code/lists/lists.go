package main

// SumList return the sum of all elements in the provided list.
func SumList(list []int) int {
	var sum int

	for _, n := range list {
		sum += n
	}

	return sum
}

// SumListsInPlace returns a list that contains the sum of all elements in the provided lists.
func SumListsInPlace(lists ...[]int) []int {
	output := []int{}

	for _, list := range lists {
		sum := SumList(list)
		output = append(output, sum)
	}

	return output
}

// SumTails return a lists that contains the sums of "tails" of all provided lists.
func SumTails(lists ...[]int) []int {
	output := []int{}

	for _, list := range lists {
		var sum int

		if len(list) == 0 {
			sum = 0
		} else {
			tail := list[1:]
			sum = SumList(tail)
		}

		output = append(output, sum)
	}

	return output
}

func main() {
	//
}
