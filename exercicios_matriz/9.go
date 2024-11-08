package exercicios_matriz

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func construirMatrizFibonacci(matriz [][]int, n, i, j int) {
	if i >= n {
		return
	}

	indiceFibonacci := i*n + j
	matriz[i][j] = fibonacci(indiceFibonacci)

	if j < n-1 {
		construirMatrizFibonacci(matriz, n, i, j+1)
	} else {
		construirMatrizFibonacci(matriz, n, i+1, 0)
	}
}

func MatrizFibo() {
	n := 3

	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, n)
	}

	construirMatrizFibonacci(matriz, n, 0, 0)

	fmt.Println("Matriz de Fibonacci:")
	for _, linha := range matriz {
		fmt.Println(linha)
	}
}
