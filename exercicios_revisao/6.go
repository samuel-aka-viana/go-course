package exercicios_revisao

import "fmt"

func SliceByIndex() {
	originalSlice := []int{10, 20, 30, 40, 50, 60}

	newSlice := originalSlice[3:]
	fmt.Println("Novo slice contendo os três últimos elementos:")
	fmt.Println(newSlice)
}
