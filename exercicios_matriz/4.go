package exercicios_matriz

import "fmt"

func somaDiagonalPrincipal(matriz [3][3]int) int {
	soma := 0

	for i := 0; i < 3; i++ {
		soma += matriz[i][i]
	}

	return soma
}

func DiagonalPrincipal() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	resultado := somaDiagonalPrincipal(matriz)
	fmt.Println("Soma dos elementos da diagonal principal:", resultado)
}
