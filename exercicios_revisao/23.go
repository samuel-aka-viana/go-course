package exercicios_revisao

import "fmt"

func copiarSlice(slice []int) []int {
	copia := append([]int(nil), slice...)
	return copia
}

func CopySlice() {
	original := []int{1, 2, 3, 4, 5}

	copia := copiarSlice(original)

	original[0] = 10

	fmt.Println("Slice original após modificação:", original)
	fmt.Println("Cópia do slice (não deve ser afetada):", copia)
}
