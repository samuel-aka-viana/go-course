package exercicios_struct

import "fmt"

//Crie uma struct chamada Pessoa com os campos nome
//(string), idade (int) e altura (float64). Em seguida,
//inicialize uma variável do tipo Pessoa e imprima
//os valores de cada campo

type Person struct {
	Nome   string
	Idade  int
	Altura float64
}

func PersonFunc() {
	pessoa := Person{Nome: "Samuel", Idade: 26, Altura: 1.86}
	fmt.Printf("Meu nome é :%s tenho:%d anos e a altura de: %.2f\n", pessoa.Nome, pessoa.Idade, pessoa.Altura)
}
