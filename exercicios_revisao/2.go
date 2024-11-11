package exercicios_revisao

import "fmt"

func AlterarIndice() {
	arr := [5]int{10, 20, 30, 40, 50}

	arr[1] = 99

	// Imprime o array após a alteração
	fmt.Println("Array após a alteração:")
	fmt.Println(arr)
}
