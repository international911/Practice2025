package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159

	radius := 5.0
	height := 10.0

	surfaceArea := 2*pi*math.Pow(radius, 2) + 2*pi*radius*height
	fmt.Printf("Полная площадь поверхности цилиндра (r=%.2f, h=%.2f) = %.2f кв. ед.\n", radius, height, surfaceArea)

	volume := pi * math.Pow(radius, 2) * height
	fmt.Printf("Объём цилиндра (r=%.2f, h=%.2f) = %.2f куб. ед.\n", radius, height, volume)
}
