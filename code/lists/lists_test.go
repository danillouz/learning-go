package main

import (
	"fmt"
	"testing"
)

func TestSumList(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := SumList(numbers)
	want := 6

	if got != want {
		t.Errorf("Got '%d', want '%d', input '%v'", got, want, numbers)
	}
}

func ExampleSumList() {
	numbers := []int{1, 1, 1, 1}
	sum := SumList(numbers)
	fmt.Println(sum)
	// Output: 4
}

func BenchmarkSumList(b *testing.B) {
	numbers := []int{4, 4, 4, 4, 4}

	for i := 0; i < b.N; i++ {
		SumList(numbers)
	}
}
