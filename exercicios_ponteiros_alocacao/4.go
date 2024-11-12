package exercicios_ponteiros_alocacao

import "fmt"

func adicionarElemento(slice *[]int, valor int) {
	*slice = append(*slice, valor)
}

func AddElement() {
	numeros := []int{1, 2, 3, 4}

	fmt.Println("Antes de adicionar:", numeros)

	adicionarElemento(&numeros, 4)

	fmt.Println("Depois de adicionar:", numeros)
}
