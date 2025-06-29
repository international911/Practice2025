package main

import "fmt"

func longestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}
	longest := strings[0]
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	result := longestString("auto", "boat", "airplane", "helicopter")
	fmt.Printf("Самая длинная строка: %q\n", result)
}
