package shapeoperation

import "math"

type Shape interface {
	Area() float64
}

// Rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle shape
type Circle struct {
	Radius float64
}

// Perimeter of circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * (c.Radius)
}

// Perimeter of rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area finds from weight and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area finds from weight and height
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
