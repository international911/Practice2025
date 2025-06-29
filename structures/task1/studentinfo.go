package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func printStudent(s Student) {
	fmt.Printf("Имя: %s, Возраст: %d, Курс: %d, Средний балл: %.2f\n", s.Name, s.Age, s.Course, s.AvgGrade)
}

func updateGrade(s *Student, newGrade float64) {
	s.AvgGrade = newGrade
	fmt.Printf("Обновлён средний балл для %s: %.2f\n", s.Name, s.AvgGrade)
}

func main() {
	student := Student{
		Name:     "Миша",
		Age:      19,
		Course:   2,
		AvgGrade: 3.6,
	}

	printStudent(student)
	updateGrade(&student, 4)
	printStudent(student)
}
