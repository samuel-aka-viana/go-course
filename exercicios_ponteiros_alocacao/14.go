package exercicios_ponteiros_alocacao

import "fmt"

func alocarMatriz(n, m int) *[][]int {
	matriz := make([][]int, n)

	for i := range matriz {
		matriz[i] = make([]int, m)
	}

	return &matriz
}

func AllocMatriz() {
	n, m := 3, 4

	matriz := alocarMatriz(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			(*matriz)[i][j] = i + j
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", (*matriz)[i][j])
		}
		fmt.Println()
	}
}
