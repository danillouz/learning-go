package rectangle

// Rectangle is a shape that has a Width and Height.
type Rectangle struct {
	Width, Height float64
}

// Perimeter returns the perimeter of a Rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a Rectangle.
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
