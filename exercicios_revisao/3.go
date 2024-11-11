package exercicios_revisao

import "fmt"

func contains(arr [5]int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func ArrayContains() {
	arr := [5]int{10, 20, 30, 40, 50}
	value := 30

	if contains(arr, value) {
		fmt.Printf("O valor %d está presente no array.\n", value)
	} else {
		fmt.Printf("O valor %d não está presente no array.\n", value)
	}
}
