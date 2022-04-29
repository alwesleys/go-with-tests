package fundamentals

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

type Triangle struct {
	Base   float64
	Height float64
}

type Shapes interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 20.0})
	var want float64 = 60.0

	if got != want {
		t.Errorf("got '%f' want '%f'", got, want)
	}
}

func TestArea(t *testing.T) {

	shapeToTest := []struct {
		shape Shapes
		want  float64
	}{
		{shape: Rectangle{10.0, 20.0}, want: 200.0},
		{shape: Circle{10.0}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, s := range shapeToTest {
		checkArea(t, s.shape, s.want)
	}
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func checkArea(t *testing.T, s Shapes, want float64) {
	got := s.Area()
	if got != want {
		t.Errorf("%#v got %g want %g", s, got, want)
	}
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
