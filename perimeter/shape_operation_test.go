package shapeoperation

import (
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

	areaTestCases := []struct {
		shape Shape
		want  float64
	}{
		{
			shape: Rectangle{5, 6},
			want:  30.0,
		},
		{
			shape: Circle{10},
			want:  314.1592653589793,
		},
		{
			shape: Triangle{10, 10},
			want:  50.0,
		},
	}

	for _, tc := range areaTestCases {
		got := tc.shape.Area()
		if got != tc.want {
			t.Errorf("got %.2f want %.2f", got, tc.want)
		}
	}
}
