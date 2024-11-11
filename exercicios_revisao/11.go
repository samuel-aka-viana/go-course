package exercicios_revisao

import "fmt"

func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func BubbleSlice() {
	slice := []int{50, 20, 40, 10, 30}

	bubbleSort(slice)

	fmt.Println("Slice ordenado em ordem crescente:")
	fmt.Println(slice)
}
