package exercicios_ponteiros_alocacao

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func ChangeVariables() {
	initValue := 20
	finalValue := 40
	fmt.Printf("%d,%d antes da troca\n", initValue, finalValue)
	swap(&initValue, &finalValue)
	fmt.Printf("%d,%d apos da troca\n", initValue, finalValue)
}
