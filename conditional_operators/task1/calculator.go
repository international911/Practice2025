package main

import (
	"fmt"
)

func calculate(a, b float64, operator string) (float64, bool) {
	switch operator {
	case "+":
		return a + b, true
	case "-":
		return a - b, true
	case "*":
		return a * b, true
	case "/":
		if b == 0 {
			return 0, false
		}
		return a / b, true
	default:
		return 0, false
	}
}

func main() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	result, ok := calculate(a, b, operator)
	if ok {
		fmt.Printf("%.2f %s %.2f = %.2f\n", a, operator, b, result)
	} else {
		fmt.Println("Ошибка: неверный оператор или деление на ноль")
	}
}
