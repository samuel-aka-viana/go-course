package exercicios_struct

// Adicione um campo cidade (string) à struct Pessoa.
//Atualize o slice do exercício 2 para incluir a cidade de cada pessoa.
//Em seguida, imprima o nome e a cidade de cada pessoa
import "fmt"

type NewPerson struct {
	Name   string
	Idade  int
	Altura float64
	Cidade string
}

func PersonCity() {
	person := NewPerson{Name: "Gilberto", Idade: 45, Altura: 1.80, Cidade: "Manaus"}
	fmt.Printf("Meu nome e %s tenho %d anos, minha altura %.2f e moro em %s\n", person.Name, person.Idade, person.Altura, person.Cidade)
}
