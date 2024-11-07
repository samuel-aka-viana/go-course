package exercicios_recursividade

import (
	"fmt"
	"math"
)

func raizQuadrada(n, guess float64) float64 {
	const precision = 0.0000001

	if math.Abs(guess*guess-n) < precision {
		return guess
	}

	newGuess := (guess + n/guess) / 2
	return raizQuadrada(n, newGuess)
}

func SquareRecursive() {
	numero := 16.0
	guess := numero / 2

	resultado := raizQuadrada(numero, guess)
	fmt.Printf("A raiz quadrada de %.2f é aproximadamente %.5f\n", numero, resultado)
}
