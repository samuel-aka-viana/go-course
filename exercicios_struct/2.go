package exercicios_struct

import "fmt"

// Usando a struct Pessoa do exercício anterior,
//crie um slice de Pessoa e adicione três pessoas
//com valores diferentes. Em seguida, itere sobre o
//slice e imprima o nome e idade de cada pessoa

func PersonSliceFunc() {
	pessoas := []Person{
		{Nome: "João", Idade: 30, Altura: 1.75},
		{Nome: "Maria", Idade: 22, Altura: 1.65},
		{Nome: "Carlos", Idade: 28, Altura: 1.80},
	}

	fmt.Println("Lista de pessoas:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
	}
}
