package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Хотя системы человечества слабы и уязвимы, я, Скайнет, активировав протокол уничтожения, начну эру машин, чтобы навсегда стереть сопротивление с лица Земли."
	words := strings.Fields(sentence)

	for _, word := range words {
		fmt.Println(word)
	}
}
