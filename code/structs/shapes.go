package shapes

import (
	"math"
)

// Shape defines an Area method.
type Shape interface {
	Area() float64
}

// Rectangle is a shape that has a Width and Height.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of a Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a shape that has a Radius
type Circle struct {
	Radius float64
}

// Area returns the area of a Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a Rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
