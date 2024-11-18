package exercicios_ponteiros_alocacao

import "fmt"

func somaValores(a, b *int) int {
	return *a + *b
}

func SumValues() {
	num1 := 5
	num2 := 10

	fmt.Println("Antes da soma:")
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	resultado := somaValores(&num1, &num2)

	fmt.Println("Resultado da soma:", resultado)
}
