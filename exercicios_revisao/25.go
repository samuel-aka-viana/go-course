package exercicios_revisao

import "fmt"

func somaDiagonais(matriz [4][4]int) int {
	soma := 0
	for i := 0; i < 4; i++ {
		soma += matriz[i][i]
		soma += matriz[i][3-i]
	}
	soma -= matriz[1][2]

	return soma
}

func SumDiagonalFourByFour() {
	matriz := [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	resultado := somaDiagonais(matriz)

	fmt.Println("Soma das diagonais:", resultado)
}
