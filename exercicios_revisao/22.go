package exercicios_revisao

import "fmt"

func saoMatrizesIguais(matriz1, matriz2 [2][2]int) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if matriz1[i][j] != matriz2[i][j] {
				return false
			}
		}
	}
	return true
}

func EqualMatriz() {
	matriz1 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	matriz2 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	matriz3 := [2][2]int{
		{1, 2},
		{5, 6},
	}

	fmt.Println("Matriz 1 e Matriz 2 são iguais?", saoMatrizesIguais(matriz1, matriz2))
	fmt.Println("Matriz 1 e Matriz 3 são iguais?", saoMatrizesIguais(matriz1, matriz3))
}
