package exercicios_recursividade

import (
	"fmt"
)

func contarOcorrencias(s string, char rune) int {
	if len(s) == 0 {
		return 0
	}
	firstChar := rune(s[0])
	if firstChar == char {
		return 1 + contarOcorrencias(s[1:], char)
	}
	return contarOcorrencias(s[1:], char)
}

func OcurrenciesCount() {
	fmt.Println(contarOcorrencias("banana", 'a')) // Output: 3
}
