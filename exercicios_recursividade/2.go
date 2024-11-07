package exercicios_recursividade

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Fatorial() {
	numero := 5
	fmt.Printf("Fatorial de %d e %d\n", numero, factorial(numero))
}
