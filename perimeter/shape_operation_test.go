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
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{20.0, 30.0}
		got := rectangle.Area()
		want := 600.0

		if got != want {
			t.Errorf(" got %0.2f want %0.2f", got, want)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * 100

		if got != want {
			t.Errorf(" got %0.2f want %0.2f", got, want)
		}
	})

}
