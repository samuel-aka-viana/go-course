package exercicios_ponteiros_alocacao

import "fmt"

func DoublePonteiro() {
	n := 5
	ponteiros := make([]*int, n)

	for i := 0; i < n; i++ {
		valor := i + 1
		ponteiros[i] = &valor
	}

	fmt.Println("Valores originais:")
	for i := 0; i < n; i++ {
		fmt.Printf("ponteiros[%d] = %d\n", i, *ponteiros[i])
	}

	for i := 0; i < n; i++ {
		*ponteiros[i] = *ponteiros[i] * 2
	}

	fmt.Println("\nValores após dobrar:")
	for i := 0; i < n; i++ {
		fmt.Printf("ponteiros[%d] = %d\n", i, *ponteiros[i])
	}
}
