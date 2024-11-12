package exercicios_ponteiros_alocacao

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
}

type PersonP struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func alterarNumeroEndereco(pessoa *PersonP, novoNumero int) {
	pessoa.Endereco.Numero = novoNumero
}

func EnderecoPointer() {
	pessoa := PersonP{
		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua ABC",
			Numero: 100,
		},
	}

	fmt.Printf("Antes da alteração: %+v\n", pessoa)

	alterarNumeroEndereco(&pessoa, 200)

	fmt.Printf("Após da alteração: %+v\n", pessoa)
}
