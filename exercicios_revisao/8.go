package exercicios_revisao

import "fmt"

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveByIndex() {
	slice := []int{10, 20, 30, 40, 50}

	newSlice := removeElement(slice, 2)

	fmt.Println("Novo slice após remover o elemento na posição 2:")
	fmt.Println(newSlice)
}
