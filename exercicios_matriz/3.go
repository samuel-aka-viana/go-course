package exercicios_matriz

import "fmt"

func produtoMatrizes(matrizA, matrizB [3][3]int) [3][3]int {
	var resultado [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			resultado[i][j] = 0
			for k := 0; k < 3; k++ {
				resultado[i][j] += matrizA[i][k] * matrizB[k][j]
			}
		}
	}

	return resultado
}

func ProdutoMatriz() {
	matrizA := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrizB := [3][3]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	resultado := produtoMatrizes(matrizA, matrizB)
	fmt.Println("Resultado do produto das matrizes:")
	for _, linha := range resultado {
		fmt.Println(linha)
	}
}
