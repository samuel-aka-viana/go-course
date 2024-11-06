package exercicio_slice

import "fmt"

func ListPersonByGender() {
	nomes := []string{"Ana", "Carlos", "Bianca", "Pedro", "Mariana"}

	sexos := []string{"F", "M", "F", "M", "F"}

	for i := 0; i < len(nomes); i++ {
		fmt.Printf("Nome: %s, Sexo: %s\n", nomes[i], sexos[i])
	}
}
