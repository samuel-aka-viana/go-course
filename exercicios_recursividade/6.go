package exercicios_recursividade

import "fmt"

func potencia(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * potencia(a, b-1)
}

func PotenciaRecursivo() {
	fmt.Println(potencia(2, 3)) // Output: 8
}
