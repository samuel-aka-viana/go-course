package exercicios_revisao

import "fmt"

func countEvenNumbers(slice []int) int {
	count := 0
	for _, num := range slice {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

func EvenCount() {
	slice := []int{10, 15, 20, 25, 30, 35, 40}

	evenCount := countEvenNumbers(slice)

	fmt.Printf("Número de números pares no slice: %d\n", evenCount)
}
