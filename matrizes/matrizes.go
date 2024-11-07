package matrizes

import "fmt"

func Matrizes() {
	// criando matriz
	fmt.Println("Criar Matrizes")
	var matrix [2][3]int
	fmt.Println(matrix)

	fmt.Println("Atribuir matriz")
	// atribuindo matriz
	nestedMatrix := [2][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}

	fmt.Println(nestedMatrix)
	fmt.Println("Percorrer matriz Indice")

	//	percorrendo matriz
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(nestedMatrix[i][j])
		}
	}
	fmt.Println("Percorrer matriz range")

	for i, row := range nestedMatrix {
		for j, value := range row {
			fmt.Printf("Element [%d][%d]=[%d]\n", i, j, value)
		}
	}
	fmt.Println("Mostrar diagonal principal")

	// mostrando diagonal princial
	matrizPrincial := [3][3]int{
		{1, 2, 3},
		{3, 4, 5},
		{5, 6, 7},
	}
	fmt.Println(matrizPrincial)

	for i, row := range matrizPrincial {
		for j, value := range row {
			if i == j {
				fmt.Printf("Diagonal princial %d\n", value)
			}
		}
	}

	fmt.Println("\nDiagonal Secundária:")
	for i, row := range matrizPrincial {
		for j, value := range row {
			if i+j == len(matrizPrincial)-1 {
				fmt.Printf("Diagonal Secundária %d\n", value)
			}
		}
	}
}
