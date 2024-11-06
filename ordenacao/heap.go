package ordenacao

import (
	"fmt"
	"math/rand"
	"time"
)

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func HeapSort() {
	rand.Seed(time.Now().UnixNano())

	arr := []int{
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
	}

	fmt.Println("Array original:", arr)
	heapSort(arr)
	fmt.Println("Array ordenado:", arr)
}
