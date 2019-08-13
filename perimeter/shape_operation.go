package shapeoperation

// Rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter finds from weight and height
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area finds from weight and height
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
