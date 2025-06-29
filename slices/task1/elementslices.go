package main

import "fmt"

func main() {
	myslice := make([]string, 0, 5)
	myslice = append(myslice, "Первый")
	myslice = append(myslice, "Второй")
	myslice = append(myslice, "Третий")

	fmt.Println("Элементы среза:")
	for i, item := range myslice {
		fmt.Printf("%d: %s\n", i, item)
	}
}
