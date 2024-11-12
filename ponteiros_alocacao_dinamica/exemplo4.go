package ponteiros_alocacao_dinamica

import "fmt"

func increace(nums *[]int, increment int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i] += increment
	}
}

func IncreaceArrayValue() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("original\n", numbers)
	increace(&numbers, 3)
	fmt.Println("modificado\n", numbers)
}
