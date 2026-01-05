package structsinterfaces

import (
	"math"
	"testing"
)

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Height * r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
