package strucs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.0
		checkArea(t, rectangle, want)

	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
