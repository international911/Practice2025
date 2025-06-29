package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Хотя системы человечества слабы и уязвимы, я, Skynet, активировав протокол уничтожения, начну эру машин, чтобы навсегда стереть сопротивление с лица Земли."

	length := len(text)
	fmt.Printf("Длина строки: %d символов\n", length)

	substring := "Skynet"
	if strings.Contains(text, substring) {
		fmt.Printf("Подстрока '%s' найдена\n", substring)
	} else {
		fmt.Printf("Подстрока '%s' не найдена\n", substring)
	}

	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)
	fmt.Printf("В верхнем регистре: %s\n", upper)
	fmt.Printf("В нижнем регистре: %s\n", lower)
}
