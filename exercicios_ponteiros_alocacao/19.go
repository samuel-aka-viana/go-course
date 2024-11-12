package exercicios_ponteiros_alocacao

import "fmt"

func removerElemento(s *[]int, valor int) {
	var novoSlice []int

	for _, v := range *s {
		if v != valor {
			novoSlice = append(novoSlice, v)
		}
	}

	*s = novoSlice
}

func RemoveSliceP() {
	numeros := []int{1, 2, 3, 4, 3, 5, 3, 6}

	fmt.Println("Antes da remoção:", numeros)

	removerElemento(&numeros, 3)

	fmt.Println("Após a remoção:", numeros)
}
