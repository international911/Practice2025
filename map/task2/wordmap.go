package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	frequency := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func main() {
	text := "Хотя системы человечества слабы и уязвимы, я, Скайнет, активировав протокол уничтожения, начну эру машин, чтобы навсегда стереть сопротивление с лица Земли."
	freq := wordFrequency(text)

	fmt.Println("Частота слов:")
	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
