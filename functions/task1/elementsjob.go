package main

import (
	"fmt"
	"sort"
)

func findElement(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func sortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func filterSlice(slice []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{100, 13, 52, 99, 7}

	index, found := findElement(numbers, 13)
	if found {
		fmt.Printf("Элемент найден на индекс %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	sorted := sortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	even := filterSlice(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Чётные числа:", even)
}
