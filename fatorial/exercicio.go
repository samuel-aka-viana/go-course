package fatorial

import "fmt"

func transformarNumero(n int) int {
	if n == 1 {
		return 1
	}

	if n%2 == 0 {
		fmt.Printf("%d é par, dividindo por 2: ", n)
		return transformarNumero(n / 2)
	}

	fmt.Printf("%d é ímpar, somando 1: ", n)
	return transformarNumero(n + 1)
}

func Exercicio() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	resultado := transformarNumero(numero)
	fmt.Printf("O número transformado chegou a: %d\n", resultado)
}
