package rectangle

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{15.0, 15.0}
	got := Perimeter(rect)
	want := 60.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func BenchmarkPerimeter(b *testing.B) {
	rect := Rectangle{5.0, 7.0}

	for i := 0; i < b.N; i++ {
		Perimeter(rect)
	}
}

func ExamplePerimeter() {
	rect := Rectangle{7.0, 4.0}
	perimeter := Perimeter(rect)
	fmt.Println(perimeter)
	// Output: 22
}

func TestArea(t *testing.T) {
	rect := Rectangle{5.0, 5.0}
	got := Area(rect)
	want := 25.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func BenchmarkArea(b *testing.B) {
	rect := Rectangle{11.0, 56.0}

	for i := 0; i < b.N; i++ {
		Area(rect)
	}
}

func ExampleArea() {
	rect := Rectangle{4.0, 4.0}
	area := Area(rect)
	fmt.Println(area)
	// Output: 16
}
