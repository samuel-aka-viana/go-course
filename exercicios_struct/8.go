package exercicios_struct

import "fmt"

// Adicione um método
// Apresentacao à struct Pessoa
// que retorne uma string com o nome e a idade formatados,
// por exemplo: "Olá, meu nome é João e eu tenho 30 anos.". Crie uma pessoa, chame o método e imprima o resultado
func (p Person) Apresentacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

func PersonShow() {
	pessoa := Person{Nome: "João", Idade: 30, Altura: 1.75}

	fmt.Println(pessoa.Apresentacao())
}
