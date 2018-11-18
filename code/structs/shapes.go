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

// Triangle is a shape that has a Height and Base.
type Triangle struct {
	Height, Base float64
}

// Area returns the area of a Triangle.
func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
