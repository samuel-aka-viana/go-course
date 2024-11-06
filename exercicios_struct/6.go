package exercicios_struct

import "fmt"

//Crie um map onde a chave é o nome da pessoa (string)
//e o valor é do tipo Pessoa. Adicione pelo menos três
//pessoas ao map e depois itere sobre ele,
//imprimindo o nome e a idade de cada pessoa

func PersonMapFunc() {
	pessoas := map[string]Person{
		"João":   {Nome: "João", Idade: 30, Altura: 1.75},
		"Maria":  {Nome: "Maria", Idade: 22, Altura: 1.65},
		"Carlos": {Nome: "Carlos", Idade: 28, Altura: 1.80},
	}

	fmt.Println("Pessoas no map:")
	for nome, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, pessoa.Idade)
	}

}
