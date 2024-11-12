package exercicios_ponteiros_alocacao

import "fmt"

func fatorialPtr(n *int) {
	if *n <= 1 {
		return
	}
	original := *n
	*n--
	fatorialPtr(n)
	*n *= original
}

func PointerFatorial() {
	numero := 5

	fmt.Println("Antes do cálculo:", numero)

	fatorialPtr(&numero)

	fmt.Println("Resultado após cálculo do fatorial:", numero)
}
