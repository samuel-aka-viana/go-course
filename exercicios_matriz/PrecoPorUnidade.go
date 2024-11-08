package exercicios_matriz

import "fmt"

type Produto struct {
	Nome   string
	Precos [3]float64
}

var tabelaDePrecos = map[string]Produto{
	"Produto 1": {"Produto 1", [3]float64{10.0, 9.5, 9.0}},
	"Produto 2": {"Produto 2", [3]float64{20.0, 19.5, 19.0}},
	"Produto 3": {"Produto 3", [3]float64{30.0, 29.5, 29.0}},
}

func calcularPreco(produtoNome string, quantidade int) float64 {
	produto, exists := tabelaDePrecos[produtoNome]
	if !exists {
		fmt.Println("Produto não encontrado!")
		return 0.0
	}

	if quantidade == 1 {
		return produto.Precos[0]
	} else if quantidade <= 5 {
		return produto.Precos[1]
	} else {
		return produto.Precos[2]
	}
}

func PrecoPorUnidade() {
	fmt.Println("Preço Produto 1, 3 unidades:", calcularPreco("Produto 1", 3))
	fmt.Println("Preço Produto 2, 7 unidades:", calcularPreco("Produto 2", 7))
	fmt.Println("Preço Produto 3, 2 unidades:", calcularPreco("Produto 3", 2))
	fmt.Println("Preço Produto 4, 3 unidades:", calcularPreco("Produto 4", 3))
}
