package exercicios_revisao

import "fmt"

func AppendSlice() {
	var slice []int

	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
	}

	fmt.Println("Slice após adicionar os números de 1 a 5:")
	fmt.Println(slice)
}
