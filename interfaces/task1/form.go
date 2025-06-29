package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 8}
	circ := Circle{Radius: 5}

	printArea(rect)
	printArea(circ)
}
