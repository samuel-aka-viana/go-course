package exercicios_ponteiros_alocacao

import "fmt"

func somaSlice(s *[]int) int {
	soma := 0
	for _, valor := range *s {
		soma += valor
	}
	return soma
}

func SumSum() {
	numeros := []int{1, 2, 3, 4, 5}

	resultado := somaSlice(&numeros)

	fmt.Printf("A soma dos elementos do slice é: %d\n", resultado)
}
