package main

import (
	"fmt"
	"unicode"
)

func countNumbers(arr []string) int {
	count := 0
	for _, char := range arr {
		if len(char) == 1 && unicode.IsDigit(rune(char[0])) {
			count++
		}
	}
	return count
}

func main() {
	examples := [][]string{
		{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"},
		{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"},
		{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"},
	}

	for i, example := range examples {
		fmt.Printf("Soal %d: %v\n", i+1, example)
		fmt.Printf("Jumlah angka: %d\n\n", countNumbers(example))
	}
}
