package exercicios_revisao

import (
	"fmt"
	"time"
)

func fatorialRecursivo(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorialRecursivo(n-1)
}

func fatorialIterativo(n int) int {
	resultado := 1
	for i := 1; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func IteractiveXRecursive() {
	numero := 20

	start := time.Now()
	resultadoRecursivo := fatorialRecursivo(numero)
	duracaoRecursivo := time.Since(start)

	start = time.Now()
	resultadoIterativo := fatorialIterativo(numero)
	duracaoIterativo := time.Since(start)

	fmt.Printf("Fatorial recursivo de %d: %d\n", numero, resultadoRecursivo)
	fmt.Printf("Tempo (recursivo): %s\n", duracaoRecursivo)

	fmt.Printf("Fatorial iterativo de %d: %d\n", numero, resultadoIterativo)
	fmt.Printf("Tempo (iterativo): %s\n", duracaoIterativo)
}
