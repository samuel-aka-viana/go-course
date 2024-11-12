package exercicios_ponteiros_alocacao

import "fmt"

func incrementarSalario(m *map[string]float64, taxa float64) {
	for chave, salario := range *m {
		(*m)[chave] = salario * (1 + taxa/100)
	}
}

func IncreaseSalary() {
	salarios := map[string]float64{
		"Alice": 3000,
		"Bob":   2500,
		"Carol": 4000,
	}

	fmt.Println("Salários antes do incremento:", salarios)

	incrementarSalario(&salarios, 10)

	fmt.Println("Salários após o incremento:", salarios)
}
