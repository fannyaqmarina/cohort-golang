package main

import (
	"fmt"
	"strings"
)

func countAlphabets(word string) map[rune]int {
	counts := make(map[rune]int)
	word = strings.ToLower(word)
	for _, char := range word {
		fmt.Printf("%c \n", char)
		if char >= 'a' && char <= 'z' {
			counts[char]++
		}
	}

	return counts
}

func main() {
	word := "selamat malam"
	alphabetCounts := countAlphabets(word)

	for char, count := range alphabetCounts {
		fmt.Printf("%c: %d | ", char, count)
	}
}
