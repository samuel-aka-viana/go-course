package ordenacao

import (
	"fmt"
)

type Aluno struct {
	Nome      string
	Matricula string
	Notas     []float64
	Media     float64
}

func calcularMedia(a Aluno) Aluno {
	var soma float64
	for _, nota := range a.Notas {
		soma += nota
	}
	a.Media = soma / float64(len(a.Notas))
	return a
}

func bubbleSortAlunos(alunos []Aluno) {
	n := len(alunos)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if alunos[j].Media < alunos[j+1].Media {
				alunos[j], alunos[j+1] = alunos[j+1], alunos[j]
			}
		}
	}
}

func BubbleAlunos() {
	alunos := []Aluno{
		{Nome: "João", Matricula: "12345", Notas: []float64{7.5, 8.0, 6.5, 9.0}},
		{Nome: "Maria", Matricula: "67890", Notas: []float64{5.5, 6.0, 7.0, 8.5}},
		{Nome: "Carlos", Matricula: "11223", Notas: []float64{9.0, 9.5, 8.5, 10.0}},
		{Nome: "Ana", Matricula: "44556", Notas: []float64{6.0, 6.5, 7.0, 7.5}},
	}

	for i := range alunos {
		alunos[i] = calcularMedia(alunos[i])
	}

	bubbleSortAlunos(alunos)

	fmt.Println("Alunos ordenados por média:")
	for _, aluno := range alunos {
		fmt.Printf("Nome: %s, Matrícula: %s, Média: %.2f\n", aluno.Nome, aluno.Matricula, aluno.Media)
	}
}
