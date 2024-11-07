package exercicios_recursividade

import "fmt"

func isPalindromo(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindromo(s[1 : len(s)-1])
}

func Palindromo() {
	strings := []string{"arara", "golang", "radar", "abacaxi"}

	for _, str := range strings {
		if isPalindromo(str) {
			fmt.Printf("\"%s\" é um palíndromo!\n", str)
		} else {
			fmt.Printf("\"%s\" não é um palíndromo.\n", str)
		}
	}
}
