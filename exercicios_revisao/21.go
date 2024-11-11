package exercicios_revisao

import "fmt"

func imprimirMatriz(matriz [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
}

func ShowMatriz() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	imprimirMatriz(matriz)
}
