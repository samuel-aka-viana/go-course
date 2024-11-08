package exercicios_matriz

import "fmt"

func somaDiagonalSecundaria(matriz [3][3]int) int {
	soma := 0

	for i := 0; i < 3; i++ {
		soma += matriz[i][2-i]
	}

	return soma
}

func DiagonalSecundaria() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	resultado := somaDiagonalSecundaria(matriz)
	fmt.Println("Soma dos elementos da diagonal secundária:", resultado)
}
