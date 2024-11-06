package ordenacao

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func MergeSort() {
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

	sortedArr := mergeSort(arr)

	fmt.Println("Array ordenado:", sortedArr)
}
