package exercicios_revisao

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

type ProdutoComDesconto struct {
	Nome     string
	Preco    float64
	Desconto float64
}

func aplicarDesconto(p Produto) ProdutoComDesconto {
	desconto := p.Preco * 0.10
	return ProdutoComDesconto{
		Nome:     p.Nome,
		Preco:    p.Preco,
		Desconto: desconto,
	}
}

func DiscountProduct() {
	produto := Produto{
		Nome:  "Camiseta",
		Preco: 100.0,
	}

	produtoComDesconto := aplicarDesconto(produto)

	fmt.Printf("Produto: %s\n", produtoComDesconto.Nome)
	fmt.Printf("Preço original: R$ %.2f\n", produtoComDesconto.Preco)
	fmt.Printf("Desconto: R$ %.2f\n", produtoComDesconto.Desconto)
	fmt.Printf("Preço com desconto: R$ %.2f\n", produtoComDesconto.Preco-produtoComDesconto.Desconto)
}
