package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{height: 10.0, width: 10.0}
	got := rectangle.Perimeter()
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
			t.Errorf("%#v, got %.2f want %.2f", shape, got, want)
		}
	}

	t.Run("area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 10.0, width: 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})

	t.Run("area of a circle", func(t *testing.T) {
		circle := Circle{radius: 10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	t.Run("area of a triangle", func(t *testing.T) {
		triangle := Triangle{base: 12, height: 6}
		want := 36.0
		checkArea(t, triangle, want)

	})
}
