package shapeoperation

// Rectangle shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter finds from weight and height
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area finds from weight and height
func Area(width, height float64) float64 {
	return width * height
}
