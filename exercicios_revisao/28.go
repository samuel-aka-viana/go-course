package exercicios_revisao

import "fmt"

func matrizTransposta(matriz [3][3]int) [3][3]int {
	var transposta [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposta[j][i] = matriz[i][j]
		}
	}

	return transposta
}

func Transposta() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	transposta := matrizTransposta(matriz)

	fmt.Println("Matriz Original:")
	for i := 0; i < 3; i++ {
		fmt.Println(matriz[i])
	}

	fmt.Println("\nMatriz Transposta:")
	for i := 0; i < 3; i++ {
		fmt.Println(transposta[i])
	}
}
