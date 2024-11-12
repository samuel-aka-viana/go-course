package exercicios_ponteiros_alocacao

import "fmt"

func aplicarFuncao(s *[]int, f func(int) int) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] = f((*s)[i])
	}
}

func TransformPointer() {
	numeros := []int{1, 2, 3, 4, 5}

	dobrar := func(x int) int {
		return x * 2
	}

	fmt.Println("Antes da transformação:", numeros)

	aplicarFuncao(&numeros, dobrar)

	fmt.Println("Após a transformação:", numeros)
}
