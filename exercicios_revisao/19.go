package exercicios_revisao

import "fmt"

func multiplicarMatrizes(m1, m2 [2][2]int) [2][2]int {
	var resultado [2][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			resultado[i][j] = m1[i][0]*m2[0][j] + m1[i][1]*m2[1][j]
		}
	}

	return resultado
}

func MatrizMult() {
	matriz1 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	matriz2 := [2][2]int{
		{5, 6},
		{7, 8},
	}

	resultado := multiplicarMatrizes(matriz1, matriz2)

	fmt.Println("Matriz Resultante:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", resultado[i][j])
		}
		fmt.Println()
	}
}
