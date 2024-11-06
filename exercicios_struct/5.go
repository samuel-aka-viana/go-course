package exercicios_struct

import "fmt"

type Aluno struct {
	Nome      string
	Matricula int
	Notas     []float64
}

func (aluno Aluno) Media() float64 {
	soma := 0.0
	for _, nota := range aluno.Notas {
		soma += nota
	}
	return soma / float64(len(aluno.Notas))
}

func (aluno Aluno) Situacao() string {
	if aluno.Media() >= 7.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func CrudStruct() {
	var alunos []Aluno

	for {
		var opcao int
		fmt.Println("\nMenu:")
		fmt.Println("1. Cadastrar aluno")
		fmt.Println("2. Listar alunos com notas e situação")
		fmt.Println("3. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 1 {
			var nome string
			var matricula int
			var notas []float64

			fmt.Print("Nome do aluno: ")
			fmt.Scan(&nome)
			fmt.Print("Matrícula do aluno: ")
			fmt.Scan(&matricula)

			fmt.Println("Digite as 4 notas do aluno:")
			for i := 0; i < 4; i++ {
				var nota float64
				fmt.Printf("Nota %d: ", i+1)
				fmt.Scan(&nota)
				notas = append(notas, nota)
			}

			alunos = append(alunos, Aluno{Nome: nome, Matricula: matricula, Notas: notas})
			fmt.Println("Aluno cadastrado com sucesso!")

		} else if opcao == 2 {
			fmt.Println("\nLista de alunos:")
			for _, aluno := range alunos {
				fmt.Printf("Nome: %s, Matrícula: %d\n", aluno.Nome, aluno.Matricula)
				fmt.Println("Notas:", aluno.Notas)
				fmt.Printf("Média: %.2f\n", aluno.Media())
				fmt.Printf("Situação: %s\n", aluno.Situacao())
				fmt.Println("----")
			}

		} else if opcao == 3 {
			fmt.Println("Saindo do programa...")
			break

		} else {
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
