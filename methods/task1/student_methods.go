package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	AvgGrade  float64
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	if s.AvgGrade >= 4.5 {
		return "Отличник"
	} else if s.AvgGrade >= 3.5 {
		return "Хорошист"
	} else {
		return "Троечник"
	}
}

func main() {
	student := Student{
		Name:      "Миша",
		BirthYear: 2006,
		AvgGrade:  3.6,
	}

	fmt.Printf("%s, Возраст: %d, Статус: %s\n", student.Name, student.GetAge(), student.GetStatus())
}
