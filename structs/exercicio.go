package structs

import (
	"fmt"
)

func Exercicio() {
	pessoas := make(map[int]struct {
		Nome string
		Nota float64
	})

	for i := 0; i < 3; i++ {
		var id int
		var nome string
		var nota float64

		fmt.Print("Digite o ID: ")
		fmt.Scan(&id)
		fmt.Print("Digite o nome: ")
		fmt.Scan(&nome)
		fmt.Print("Digite a nota: ")
		fmt.Scan(&nota)

		pessoas[id] = struct {
			Nome string
			Nota float64
		}{Nome: nome, Nota: nota}
	}

	var somaNotas float64
	for _, pessoa := range pessoas {
		somaNotas += pessoa.Nota
	}

	media := somaNotas / float64(len(pessoas))

	fmt.Printf("Média das notas: %.2f\n", media)
}
