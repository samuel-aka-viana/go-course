package exercicios_struct

import (
	"fmt"
)

type AlunoDesafio struct {
	Nome      string
	Matricula string
	Notas     []float64
}

func cadastrarAluno(nome, matricula string, notas []float64) AlunoDesafio {
	return AlunoDesafio{
		Nome:      nome,
		Matricula: matricula,
		Notas:     notas,
	}
}

func calcularMedia(aluno AlunoDesafio) float64 {
	var soma float64
	for _, nota := range aluno.Notas {
		soma += nota
	}
	return soma / float64(len(aluno.Notas))
}

func situacao(aluno AlunoDesafio) string {
	media := calcularMedia(aluno)
	if media >= 7.0 {
		return "Aprovado"
	}
	return "Reprovado"
}

func listarAlunos(alunos []AlunoDesafio) {
	for _, aluno := range alunos {
		media := calcularMedia(aluno)
		fmt.Printf("Nome: %s\n", aluno.Nome)
		fmt.Printf("Matrícula: %s\n", aluno.Matricula)
		fmt.Printf("Notas: %v\n", aluno.Notas)
		fmt.Printf("Média: %.2f\n", media)
		fmt.Printf("Situação: %s\n", situacao(aluno))
		fmt.Println()
	}
}

func ClassMean() {
	var alunos []AlunoDesafio

	alunos = append(alunos, cadastrarAluno("João Silva", "12345", []float64{7.5, 8.0, 6.5, 9.0}))
	alunos = append(alunos, cadastrarAluno("Maria Souza", "67890", []float64{5.0, 6.0, 7.5, 4.5}))
	alunos = append(alunos, cadastrarAluno("Carlos Pereira", "11223", []float64{8.0, 7.5, 9.0, 6.5}))

	listarAlunos(alunos)
}
