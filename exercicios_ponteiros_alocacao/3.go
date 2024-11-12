package exercicios_ponteiros_alocacao

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func alterarIdade(p *Pessoa, novaIdade int) {
	p.idade = novaIdade
}

func ChangeAgePerson() {
	pessoa := Pessoa{
		nome:  "Carlos",
		idade: 30,
	}

	fmt.Printf("Antes: %s tem %d anos\n", pessoa.nome, pessoa.idade)

	alterarIdade(&pessoa, 35)

	fmt.Printf("Depois: %s tem %d anos\n", pessoa.nome, pessoa.idade)
}
