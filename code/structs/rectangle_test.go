package rectangle

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(15.0, 15.0)
	want := 60.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Perimeter(5.0, 7.0)
	}
}

func ExamplePerimeter() {
	perimeter := Perimeter(7.0, 4.0)
	fmt.Println(perimeter)
	// Output: 22
}

func TestArea(t *testing.T) {
	got := Area(5.0, 5.0)
	want := 25.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Area(11.0, 56.0)
	}
}

func ExampleArea() {
	area := Area(4.0, 4.0)
	fmt.Println(area)
	// Output: 16
}
