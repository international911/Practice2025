package main

import "fmt"

func removeByIndex(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("Ошибка: неверный индекс")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	myslice := []string{"Первый", "Второй", "Третий", "Четвёртый"}
	fmt.Println("Исходный срез:", myslice)
	myslice = removeByIndex(myslice, 3)
	fmt.Println("Срез после удаления:", myslice)
}
