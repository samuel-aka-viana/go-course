package exercicios_ponteiros_alocacao

import "fmt"

func somaRecursiva(s *[]int) int {
	if len(*s) == 0 {
		return 0
	}
	novoN := (*s)[1:]
	return (*s)[0] + somaRecursiva(&novoN)
}

func SumRecursion() {
	numeros := []int{1, 2, 3, 4, 5}
	resultado := somaRecursiva(&numeros)
	fmt.Printf("A soma dos elementos do slice é: %d\n", resultado)
}
