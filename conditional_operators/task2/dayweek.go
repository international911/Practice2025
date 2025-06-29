package main

import (
	"fmt"
)

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	default:
		return "Неверный номер дня"
	}
}

func main() {
	var number int
	fmt.Print("Введите номер дня (1-7): ")
	fmt.Scan(&number)
	fmt.Println(dayOfWeek(number))
}
