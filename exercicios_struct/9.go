package exercicios_struct

import "fmt"

// Crie uma função que receba um slice de Pessoa e
// retorne um map onde as chaves são os nomes das cidades
// e os valores são a contagem de pessoas que vivem em cada cidade. Teste a função com o slice do exercício 2, incluindo as cidades no struct de Pessoa.
type PersonCount struct {
	Nome   string
	Idade  int
	Altura float64
	Cidade string
}

func PersonCountReturnMap() {
	pessoas := []PersonCount{
		{Nome: "João", Idade: 30, Altura: 1.75, Cidade: "São Paulo"},
		{Nome: "Maria", Idade: 22, Altura: 1.65, Cidade: "Rio de Janeiro"},
		{Nome: "Carlos", Idade: 28, Altura: 1.80, Cidade: "São Paulo"},
	}

	contagemPorCidade := ContarPessoasPorCidade(pessoas)
	fmt.Println("Contagem de pessoas por cidade:")
	for cidade, contagem := range contagemPorCidade {
		fmt.Printf("Cidade: %s, Contagem: %d\n", cidade, contagem)
	}
}

func ContarPessoasPorCidade(pessoas []PersonCount) map[string]int {
	contagem := make(map[string]int)
	for _, pessoa := range pessoas {
		contagem[pessoa.Cidade]++
	}
	return contagem
}
