package shapeoperation

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, s Shape, want float64) {
		area := s.Area()
		if area != want {
			t.Errorf(" got %0.2f want %0.2f", area, want)
		}
	}

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{20.0, 30.0}
		want := 600.0
		checkArea(t, rectangle, want)

	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := math.Pi * 100
		checkArea(t, circle, want)
	})

}
