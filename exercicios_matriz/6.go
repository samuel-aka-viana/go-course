package exercicios_matriz

import "fmt"

func contarOcorrencias(numeros []int) map[int]int {
	ocorrencias := make(map[int]int)

	for _, numero := range numeros {
		ocorrencias[numero]++
	}

	return ocorrencias
}

func MatrizOcorrencias() {
	numeros := []int{1, 2, 3, 2, 1, 3, 4, 2, 1, 5}

	resultado := contarOcorrencias(numeros)
	fmt.Println("Contagem de ocorrências:", resultado)
}
