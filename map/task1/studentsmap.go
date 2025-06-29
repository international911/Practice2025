package main

import "fmt"

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
	fmt.Printf("Добавлена оценка %d для %s\n", grade, name)
}

func findGrade(grades map[string]int, name string) (int, bool) {
	grade, exists := grades[name]
	if exists {
		fmt.Printf("Оценка для %s: %d\n", name, grade)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
	return grade, exists
}

func deleteGrade(grades map[string]int, name string) {
	if _, exists := grades[name]; exists {
		delete(grades, name)
		fmt.Printf("Оценка для %s удалена\n", name)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func main() {
	grades := make(map[string]int)
	addGrade(grades, "Егор", 5)
	addGrade(grades, "Тимур", 5)
	addGrade(grades, "Миша", 5)
	findGrade(grades, "Егор")
	findGrade(grades, "Миша")
	deleteGrade(grades, "Тимур")
	deleteGrade(grades, "Ваня")

	fmt.Println("Текущие оценки:", grades)
}
