package exercicios_matriz

import "fmt"

func encontrarMiniMax(matriz [][]int) int {
	miniMax := 1000000

	for _, linha := range matriz {
		maxNaLinha := linha[0]
		for _, valor := range linha {
			if valor > maxNaLinha {
				maxNaLinha = valor
			}
		}

		if maxNaLinha < miniMax {
			miniMax = maxNaLinha
		}
	}

	return miniMax
}

func miniMaxMatriz() {
	matriz := [][]int{
		{1, 5, 9, 2, 8},
		{4, 3, 7, 6, 1},
		{6, 8, 2, 5, 3},
		{9, 0, 6, 7, 4},
		{3, 4, 8, 1, 2},
	}

	miniMax := encontrarMiniMax(matriz)
	fmt.Println("O MiniMax da matriz é:", miniMax)
}
