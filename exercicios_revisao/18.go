package exercicios_revisao

import "fmt"

func potencia(base, expoente int) int {
	if expoente == 0 {
		return 1
	}
	return base * potencia(base, expoente-1)
}

func Exponential() {
	base := 2
	expoente := 3
	resultado := potencia(base, expoente)
	fmt.Printf("%d^%d = %d\n", base, expoente, resultado)
}
