package exercicios_revisao

import "fmt"

func buscaBinaria(slice []int, alvo, inicio, fim int) int {
	if inicio > fim {
		return -1
	}

	meio := (inicio + fim) / 2

	if slice[meio] == alvo {
		return meio
	}

	if alvo < slice[meio] {
		return buscaBinaria(slice, alvo, inicio, meio-1)
	}

	return buscaBinaria(slice, alvo, meio+1, fim)
}

func BinarySearch() {
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	alvo := 7

	posicao := buscaBinaria(slice, alvo, 0, len(slice)-1)

	if posicao == -1 {
		fmt.Println("Elemento não encontrado!")
	} else {
		fmt.Printf("Elemento %d encontrado na posição %d.\n", alvo, posicao)
	}
}
