package structs

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

type Endereco struct {
	Rua    string
	Cidade string
}

func (pessoa Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", pessoa.Nome, pessoa.Idade)
}

func Struc() {
	pessoa := Pessoa{Nome: "Samuel", Idade: 18, Cidade: "Manaus"}

	pessoa.Nome = "Samuel Viana"

	fmt.Println(pessoa)

	fmt.Println(pessoa.Saudacao())
}
