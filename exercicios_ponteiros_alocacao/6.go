package exercicios_ponteiros_alocacao

import "fmt"

func filtrarPares(slice *[]int) {
	var pares []int
	for _, valor := range *slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}
	*slice = pares
}

func FilterEven() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Antes de filtrar:", numeros)

	filtrarPares(&numeros)

	fmt.Println("Apenas pares:", numeros)
}
