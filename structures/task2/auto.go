package main

import "fmt"

type Engine struct {
	Type  string
	Power int
}

type Car struct {
	Model  string
	Year   int
	Engine Engine
}

func printCar(c Car) {
	fmt.Printf("Модель: %s, Год: %d, Двигатель: %s, Мощность: %d л.с.\n", c.Model, c.Year, c.Engine.Type, c.Engine.Power)
}

func main() {
	car := Car{
		Model: "Porsche 911 GT3",
		Year:  1999,
		Engine: Engine{
			Type:  "Бензиновый",
			Power: 510,
		},
	}

	printCar(car)
}
