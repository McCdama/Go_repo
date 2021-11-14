package interf

import (
	"fmt"
	"math"
)

// shape interface
type Shape interface {
	Area() float64
	Circumf() float64
}

type Square struct {
	Length float64
}
type Circle struct {
	Radius float64
}

// square methods
func (s Square) Area() float64 {
	return s.Length * s.Length
}
func (s Square) Circumf() float64 {
	return s.Length * 4
}

// circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Circumf() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.Area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.Circumf())
}
