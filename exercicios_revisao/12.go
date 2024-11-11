package exercicios_revisao

import "fmt"

func countWordFrequencies(words []string) map[string]int {
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func CountFrequencie() {
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}

	wordFrequencies := countWordFrequencies(words)

	fmt.Println("Frequência das palavras:")
	for word, count := range wordFrequencies {
		fmt.Printf("%s: %d\n", word, count)
	}
}
