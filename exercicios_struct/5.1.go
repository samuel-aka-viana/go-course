package exercicios_struct

import "fmt"

// Crie uma nova struct chamada Endereco
// com os campos rua (string), numero (int) e cidade (string).
// Modifique a struct Pessoa para incluir um campo endereco do tipo Endereco. Crie uma pessoa com valores para o endereço e imprima os dados completos dela
type Endereco struct {
	Rua    string
	Cidade string
	Estado string
}

type Usuario struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func PersonNestedAddress() {
	pessoa := Usuario{Nome: "Gilberto", Idade: 45, Endereco: Endereco{Rua: "Danilo aerosa", Cidade: "Manaus", Estado: "Amazonas"}}
	fmt.Printf("Meu nome %s minha idade %d e moro em %s no estado(a) %s\n", pessoa.Nome, pessoa.Idade, pessoa.Endereco.Cidade, pessoa.Endereco.Estado)
}
