package exercicio_array

import (
	"fmt"
)

func MenorValor(numeros []int) int {
	menor := numeros[0]
	for _, numero := range numeros {
		if numero < menor {
			menor = numero
		}
	}
	return menor
}

func MaiorValor() {
	numeros := []int{10, 5, 8, 3, 15, 7}
	menor := MenorValor(numeros)
	fmt.Println("O valor menor é:", menor)

}
