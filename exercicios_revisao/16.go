package exercicios_revisao

import "fmt"

type Enderecos struct {
	Rua    string
	Numero int
	Cidade string
}

type Pessoas struct {
	Nome     string
	Idade    int
	Endereco Enderecos
}

func MapNested() {
	pessoa := Pessoas{
		Nome:  "João Silva",
		Idade: 30,
		Endereco: Enderecos{
			Rua:    "Avenida Brasil",
			Numero: 123,
			Cidade: "São Paulo",
		},
	}

	fmt.Printf("Endereço de %s: %s, %d - %s\n", pessoa.Nome, pessoa.Endereco.Rua, pessoa.Endereco.Numero, pessoa.Endereco.Cidade)
}
