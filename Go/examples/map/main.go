package main

import (
	"fmt"
	"slices"
)

func main() {
	words := []string{"go", "map", "go", "slice"}
	countByWord := make(map[string]int, len(words))

	for _, word := range words {
		countByWord[word]++
	}

	keys := make([]string, 0, len(countByWord))
	for word := range countByWord {
		keys = append(keys, word)
	}
	slices.Sort(keys)

	for _, word := range keys {
		fmt.Printf("%s: %d\n", word, countByWord[word])
	}
}
