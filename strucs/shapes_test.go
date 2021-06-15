package strucs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36},
	}
	for _, test := range areaTests {
		xType := fmt.Sprintf("%T", test.shape)
		t.Run(xType, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %g want %g", got, test.want)
			}
		})

	}
}
