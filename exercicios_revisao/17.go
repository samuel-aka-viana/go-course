package exercicios_revisao

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func ViewPersons() {
	pessoas := make(map[int]Pessoa)

	pessoas[1] = Pessoa{Nome: "João Silva", Idade: 30}
	pessoas[2] = Pessoa{Nome: "Maria Oliveira", Idade: 25}
	pessoas[3] = Pessoa{Nome: "Pedro Santos", Idade: 40}

	for id, pessoa := range pessoas {
		fmt.Printf("ID: %d, Nome: %s, Idade: %d\n", id, pessoa.Nome, pessoa.Idade)
	}
}
