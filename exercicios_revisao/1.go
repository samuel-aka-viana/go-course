package exercicios_revisao

import "fmt"

func ArrayFixo() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i + 1
	}

	fmt.Println("Elementos do array:")
	for _, value := range arr {
		fmt.Println(value)
	}
}
