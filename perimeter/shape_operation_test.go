package shapeoperation

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{20.0, 30.0})
	want := 600.0

	if got != want {
		t.Errorf(" got %0.2f want %0.2f", got, want)
	}
}
