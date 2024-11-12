package exercicios_ponteiros_alocacao

import "fmt"

type Person struct {
	nome  string
	idade int
}

func novaPessoa(nome string, idade int) *Person {
	pessoa := &Person{
		nome:  nome,
		idade: idade,
	}
	return pessoa
}

func StructPointer() {
	pessoa := novaPessoa("Carlos", 30)

	fmt.Printf("Nome: %s, Idade: %d\n", pessoa.nome, pessoa.idade)
}
