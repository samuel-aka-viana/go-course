package recursividade

import "fmt"

func somaRecursiva(n int) int {
	if n == 1 {
		return 1
	}
	return n + somaRecursiva(n-1)
}

func SomaMesmoNumero() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	resultado := somaRecursiva(numero)
	fmt.Printf("A soma dos números de 1 até %d é: %d\n", numero, resultado)
}
