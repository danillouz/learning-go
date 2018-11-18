package main

import (
	"fmt"
	"reflect"
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

func TestSumListsInPlace(t *testing.T) {
	a := []int{1, 2, 1}
	b := []int{2, 2, 6}
	got := SumListsInPlace(a, b)
	want := []int{4, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got '%v', want '%v', input '%v', '%v'", got, want, a, b)
	}
}

func ExampleSumListsInPlace() {
	a := []int{4, 4}
	b := []int{2, 2}
	summed := SumListsInPlace(a, b)
	fmt.Println(summed)
	// Output: [8 4]
}

func BenchmarkSumListsInPlace(b *testing.B) {
	x := []int{3, 3}
	y := []int{5, 5}

	for i := 0; i < b.N; i++ {
		SumListsInPlace(x, y)
	}
}
