package exercicios_ponteiros_alocacao

import "fmt"

type Produto struct {
	nome  string
	preco float64
}

func aplicarDesconto(produtos *[]Produto) {
	for i := range *produtos {
		(*produtos)[i].preco = (*produtos)[i].preco * 0.9
	}
}

func DiscountPointer() {
	produtos := []Produto{
		{"Produto A", 100.0},
		{"Produto B", 150.0},
		{"Produto C", 200.0},
	}

	fmt.Println("Produtos antes do desconto:")
	for _, p := range produtos {
		fmt.Printf("%s: %.2f\n", p.nome, p.preco)
	}

	aplicarDesconto(&produtos)

	fmt.Println("\nProdutos após o desconto:")
	for _, p := range produtos {
		fmt.Printf("%s: %.2f\n", p.nome, p.preco)
	}
}
