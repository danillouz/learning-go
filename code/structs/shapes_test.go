package shapes

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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("Got %.2f, want %.2f", got, want)
		}
	}

	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		want := 25.0

		checkArea(t, rectangle, want)
	})

	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{4.0}
		want := 50.26548245743669

		checkArea(t, circle, want)
	})
}

func BenchmarkArea(b *testing.B) {
	b.Run("Rectangle area", func(b *testing.B) {
		rectangle := Rectangle{11.0, 56.0}

		for i := 0; i < b.N; i++ {
			rectangle.Area()
		}
	})

	b.Run("Circle area", func(b *testing.B) {
		circle := Circle{11.0}

		for i := 0; i < b.N; i++ {
			circle.Area()
		}
	})
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{4.0, 4.0}
	area := rectangle.Area()
	fmt.Println(area)
	// Output: 16
}

func ExampleCircle_Area() {
	circle := Circle{4.0}
	area := circle.Area()
	fmt.Println(area)
	// Output: 50.26548245743669
}
