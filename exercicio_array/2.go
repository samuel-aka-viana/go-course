package exercicio_array

import "fmt"

func SomaElementos(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaArray() {
	numeros := []int{1, 2, 3, 4, 5}
	resultado := SomaElementos(numeros)
	fmt.Println("A soma dos elementos é:", resultado)
}
