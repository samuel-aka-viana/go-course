package exercicios_revisao

import "fmt"

func bubbleSortInvertido(arr [10]int) [10]int {
	slice := arr[:]
	n := len(slice)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return arr
}

func InvertedArray() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Imprimindo o array original
	fmt.Println("Array original:", arr)

	// Ordenando o array de forma invertida usando Bubble Sort
	arrInvertido := bubbleSortInvertido(arr)

	fmt.Println("Array original:", arrInvertido)
}
