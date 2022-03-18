package structure

import "math"

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.w + rectangle.h)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.w * rectangle.h
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.r * circle.r
}

type Rectangle struct {
	w float64
	h float64
}

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
}
