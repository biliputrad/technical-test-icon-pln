package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseEachWord(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = reverseString(word)
	}
	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	reversedWords := reverseEachWord(input)

	fmt.Printf("Kalimat terbalik: %s\n", reversedWords)
}
