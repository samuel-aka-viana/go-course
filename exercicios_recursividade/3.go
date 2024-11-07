package exercicios_recursividade

import (
	"fmt"
)

func soma(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return slice[0] + soma(slice[1:])
}

func SumSlice() {
	fmt.Println(soma([]int{1, 2, 3, 4}))
}
