package exercicios_revisao

import "fmt"

func SliceConcat() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}

	concatenatedSlice := append(slice1, slice2...)

	fmt.Println("Slice concatenado:")
	fmt.Println(concatenatedSlice)
}
