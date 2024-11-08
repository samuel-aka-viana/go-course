package exercicios_matriz

import "fmt"

type Product struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func calcularValorTotal(matrizProdutos [][]Product) float64 {
	valorTotal := 0.0

	for _, linha := range matrizProdutos {
		for _, produto := range linha {
			valorTotal += produto.Preco * float64(produto.Quantidade)
		}
	}

	return valorTotal
}

func MatrizProduto() {
	matrizProdutos := [][]Product{
		{
			{"Produto A", 10.0, 5},
			{"Produto B", 20.0, 3},
		},
		{
			{"Produto C", 15.5, 2},
			{"Produto D", 7.5, 8},
		},
	}

	total := calcularValorTotal(matrizProdutos)
	fmt.Printf("Valor total dos produtos: R$ %.2f\n", total)
}
