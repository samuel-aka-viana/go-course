package ponteiros_alocacao_dinamica

import "fmt"

func increment(n *int) *int {
	*n++
	return n
}

func Ponteiros() {
	// referenciando ponteiro
	var x int = 10
	var p *int = &x
	fmt.Println(*p)
	// alterando valor

	var value int = *p
	*p = 20
	fmt.Printf("%d valor antigo do ponteiro\n", value)
	fmt.Printf("%d novo valor do ponteiro\n", *p)

	fmt.Println(*increment(&x))
}
