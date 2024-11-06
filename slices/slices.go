package slices

import "fmt"

func Slices() {
	//criacao

	var sl []int
	sl = make([]int, 5)
	sl = []int{1, 2, 3}
	fmt.Println(sl)
	// iteracao

	slInt := []int{1, 2, 3}
	for i := 0; i < len(slInt); i++ {
		fmt.Println(slInt[i])
	}

	// key value slice

	for i, v := range slInt {
		fmt.Println("indice: %d, valor %d", i, v)
	}

	// append
	slApp := []int{1, 2, 3}
	slApp = append(slApp, 4)
	fmt.Println(slApp)

	// append lista existente
	slAppNew := []int{1, 2, 3}
	slAppNew = append(slAppNew, slApp...)
	fmt.Println(slAppNew)

	//remover item do slice
	slNew := []int{1, 2, 3, 5, 6, 7}
	index := 2
	newSl := append(slNew[:index], slNew[index+1:]...)
	fmt.Println(newSl)

}
