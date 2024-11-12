package ponteiros_alocacao_dinamica

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func ChangeVariables() {
	x := 20
	y := 40
	fmt.Printf("%d,%d antes da troca\n", x, y)
	swap(&x, &y)
	fmt.Printf("%d,%d apos da troca\n", x, y)
}
