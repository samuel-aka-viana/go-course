package exercicios_ponteiros_alocacao

import "fmt"

func somarAteN(n *int) int {
	if *n == 1 {
		return 1
	}
	novoN := *n - 1
	return *n + somarAteN(&novoN)
}

func SumN() {
	num := 5

	resultado := somarAteN(&num)
	fmt.Printf("A soma de 1 até %d é: %d\n", num, resultado)
}
