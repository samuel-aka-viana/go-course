package exercicios_ponteiros_alocacao

import "fmt"

func reverterString(s *string) {
	runes := []rune(*s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	*s = string(runes)
}

func ReverseString() {
	str := "exemplo"

	fmt.Println("Antes da reversão:", str)

	reverterString(&str)

	fmt.Println("Após a reversão:", str)
}
