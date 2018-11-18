package shapes

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: Rectangle{5.0, 5.0},
			want:  25,
		},
		{
			name:  "Circle",
			shape: Circle{4.0},
			want:  50.26548245743669,
		},
		{
			name:  "Triangle",
			shape: Triangle{6.0, 3.0},
			want:  9.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("Got %.2f, want %.2f, input %#v", got, tt.want, tt.shape)
			}
		})
	}
}

func BenchmarkArea(b *testing.B) {
	b.Run("Rectangle", func(b *testing.B) {
		rectangle := Rectangle{11.0, 56.0}

		for i := 0; i < b.N; i++ {
			rectangle.Area()
		}
	})

	b.Run("Circle", func(b *testing.B) {
		circle := Circle{11.0}

		for i := 0; i < b.N; i++ {
			circle.Area()
		}
	})

	b.Run("Triangle", func(b *testing.B) {
		triangle := Triangle{6.0, 9.0}

		for i := 0; i < b.N; i++ {
			triangle.Area()
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

func ExampleTriangle_Area() {
	triangle := Triangle{5.0, 8.0}
	area := triangle.Area()
	fmt.Println(area)
	// Output: 20
}
