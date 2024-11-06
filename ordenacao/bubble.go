package ordenacao

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSort(arr []int) {
	rand.Seed(time.Now().UnixNano())
	numeros := make([]int, 7)
	for i := 0; i < len(numeros); i++ {
		numeros[i] = rand.Intn(100)
	}

	fmt.Println("Números originais:", numeros)

	bubbleSort(numeros)

	fmt.Println("Números ordenados:", numeros)
}
