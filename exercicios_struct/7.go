package exercicios_struct

import "fmt"

// Crie uma função que receba um slice de Pessoa e
//uma idade mínima. A função deve retornar um novo slice
//contendo apenas as pessoas que têm idade maior ou igual
//à idade mínima. Teste a função com o slice do exercício 2

func PersonMinAge() {
	pessoas := []Person{
		{Nome: "João", Idade: 30, Altura: 1.75},
		{Nome: "Maria", Idade: 22, Altura: 1.65},
		{Nome: "Carlos", Idade: 28, Altura: 1.80},
	}

	fmt.Println("Lista completa de pessoas:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
	}

	idadeMinima := 25
	pessoasFiltradas := FiltrarPorIdadeMinima(pessoas, idadeMinima)

	fmt.Printf("\nPessoas com idade maior ou igual a %d:\n", idadeMinima)
	for _, pessoa := range pessoasFiltradas {
		fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
	}
}

func FiltrarPorIdadeMinima(pessoas []Person, idadeMinima int) []Person {
	var pessoasFiltradas []Person
	for _, pessoa := range pessoas {
		if pessoa.Idade >= idadeMinima {
			pessoasFiltradas = append(pessoasFiltradas, pessoa)
		}
	}
	return pessoasFiltradas
}
