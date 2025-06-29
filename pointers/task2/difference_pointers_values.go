package main

import "fmt"

func incrementByValue(n int) {
	n++
	fmt.Printf("Внутри функции (по значению): n = %d\n", n)
}

func incrementByPointer(n *int) {
	*n++
	fmt.Printf("Внутри функции (по указателю): n = %d\n", *n)
}

func main() {
	value := 715
	fmt.Printf("Исходное значение: value = %d\n", value)

	incrementByValue(value)
	fmt.Printf("После функции по значению): value = %d\n", value)

	incrementByPointer(&value)
	fmt.Printf("После функции по указателю): value = %d\n", value)
}
