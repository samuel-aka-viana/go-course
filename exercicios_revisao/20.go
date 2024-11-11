package exercicios_revisao

import "fmt"

func somaDasDiagonais(matriz [3][3]int) (int, int) {
	var somaPrincipal, somaSecundaria int

	for i := 0; i < 3; i++ {
		somaPrincipal += matriz[i][i]
	}

	for i := 0; i < 3; i++ {
		somaSecundaria += matriz[i][2-i]
	}

	return somaPrincipal, somaSecundaria
}

func SumDiagonals() {
	matriz := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	somaPrincipal, somaSecundaria := somaDasDiagonais(matriz)

	fmt.Printf("Soma da Diagonal Principal: %d\n", somaPrincipal)
	fmt.Printf("Soma da Diagonal Secundária: %d\n", somaSecundaria)
}
