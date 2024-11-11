package exercicios_revisao

import "fmt"

func ModifiedSlice() {
	slice := []int{10, 20, 30, 40, 50}

	slice[1] = 99

	fmt.Println("Slice após a modificação:")
	fmt.Println(slice)
}
