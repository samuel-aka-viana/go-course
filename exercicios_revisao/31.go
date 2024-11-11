package exercicios_revisao

import (
	"fmt"
)

type Endereco struct {
	Estado string
	Cidade string
}

type Person struct {
	Nome      string
	Sexo      string
	Idade     int
	Empregado string
	Endereco  Endereco
}

func mediaMulheresRecursiva(pessoas []Person, estado string, index int, count int, totalIdade int) (float64, int) {
	if index >= len(pessoas) {
		if count == 0 {
			return 0, count
		}
		return float64(totalIdade) / float64(count), count
	}

	pessoa := pessoas[index]
	if pessoa.Sexo == "feminino" && pessoa.Idade > 25 && pessoa.Endereco.Estado == estado {
		count++
		totalIdade += pessoa.Idade
	}

	return mediaMulheresRecursiva(pessoas, estado, index+1, count, totalIdade)
}

func contarHomensDesempregados(pessoas []Person, estado string, index int, count int) int {
	if index >= len(pessoas) {
		return count
	}

	pessoa := pessoas[index]
	if pessoa.Sexo == "masculino" && pessoa.Idade > 25 && pessoa.Empregado == "não" && pessoa.Endereco.Estado == estado {
		count++
	}

	return contarHomensDesempregados(pessoas, estado, index+1, count)
}

func main() {
	pessoas := []Person{
		{"Ana", "feminino", 30, "sim", Endereco{"São Paulo", "São Paulo"}},
		{"João", "masculino", 35, "não", Endereco{"São Paulo", "São Paulo"}},
		{"Maria", "feminino", 40, "sim", Endereco{"Minas Gerais", "Belo Horizonte"}},
		{"Carlos", "masculino", 28, "não", Endereco{"Minas Gerais", "Belo Horizonte"}},
		{"Paula", "feminino", 26, "não", Endereco{"São Paulo", "Campinas"}},
		{"Pedro", "masculino", 30, "não", Endereco{"São Paulo", "Campinas"}},
	}

	estado := "São Paulo"
	media, _ := mediaMulheresRecursiva(pessoas, estado, 0, 0, 0)
	fmt.Printf("Média de mulheres acima de 25 anos em %s: %.2f\n", estado, media)

	quantidadeHomens := contarHomensDesempregados(pessoas, estado, 0, 0)
	fmt.Printf("Quantidade de homens desempregados acima de 25 anos em %s: %d\n", estado, quantidadeHomens)
}
