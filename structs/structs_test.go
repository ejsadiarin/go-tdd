package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("with float width and height args", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("with valid args", func(t *testing.T) {
		got := Area(5.0, 5.0)
		want := 25.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
