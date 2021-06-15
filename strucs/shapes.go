package strucs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}
