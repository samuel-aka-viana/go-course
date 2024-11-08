package exercicios_matriz

import "fmt"

func transpostaMatriz(matriz [][]int) [][]int {
	n := len(matriz)
	transposta := make([][]int, n)
	for i := range transposta {
		transposta[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			transposta[i][j] = matriz[j][i]
		}
	}

	return transposta
}

func Transposta() {
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz original:")
	for _, linha := range matriz {
		fmt.Println(linha)
	}

	transposta := transpostaMatriz(matriz)
	fmt.Println("\nMatriz transposta:")
	for _, linha := range transposta {
		fmt.Println(linha)
	}
}
