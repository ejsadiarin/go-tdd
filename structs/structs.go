package structs

import (
	"math"
)

// Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
// see: https://en.wikipedia.org/wiki/Parametric_polymorphism
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
	Work   string
}

// instead of:
// func Area(r Rectangle) float64 {
// 	return r.Width * r.Height
// }

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
