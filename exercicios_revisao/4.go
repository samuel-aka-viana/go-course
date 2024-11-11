package exercicios_revisao

import "fmt"

func ArrayBySlice() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := arr[:5]

	fmt.Println("Slice contendo os primeiros 5 elementos:")
	fmt.Println(slice)
}
