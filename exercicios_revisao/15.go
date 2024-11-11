package exercicios_revisao

import "fmt"

func countGreaterThan10(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	count := 0
	if slice[0] > 10 {
		count = 1
	}

	return count + countGreaterThan10(slice[1:])
}

func GreaterThanTen() {
	numbers := []int{5, 12, 8, 15, 3, 18, 7, 20}

	result := countGreaterThan10(numbers)
	fmt.Printf("O número de elementos maiores que 10 é: %d\n", result)
}
