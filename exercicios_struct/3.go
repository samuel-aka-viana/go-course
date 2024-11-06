package exercicios_struct

import "fmt"

// Crie uma função que receba um slice de Pessoa
// e retorne a média de idade das pessoas. Use o slice criado no exercício 2
// como entrada e imprima o resultado

func PersonMean() {
	pessoas := []Person{
		{Nome: "João", Idade: 30, Altura: 1.75},
		{Nome: "Maria", Idade: 22, Altura: 1.65},
		{Nome: "Carlos", Idade: 28, Altura: 1.80},
	}

	media := CalcularMediaIdade(pessoas)
	fmt.Printf("o resultado da media de idades: %.2f", media)

}

func CalcularMediaIdade(pessoas []Person) float64 {
	if len(pessoas) == 0 {
		return 0
	}

	var somaIdade int
	for _, pessoa := range pessoas {
		somaIdade += pessoa.Idade
	}

	media := float64(somaIdade) / float64(len(pessoas))
	return media
}
