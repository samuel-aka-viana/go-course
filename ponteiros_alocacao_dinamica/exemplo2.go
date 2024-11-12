package ponteiros_alocacao_dinamica

import "fmt"

func inverted(slice *[]int) {
	s := *slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SliveInvert() {
	lista := []int{10, 20, 30, 40, 50}
	fmt.Printf("array original", lista)
	inverted(&lista)
	fmt.Printf("array invertido", lista)
}
