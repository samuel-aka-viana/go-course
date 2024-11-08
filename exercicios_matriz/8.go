package exercicios_matriz

import "fmt"

func buscaRecursiva(matriz [][]string, palavra string, i, j int) bool {
	if i >= len(matriz) {
		return false
	}

	if matriz[i][j] == palavra {
		return true
	}

	if j < len(matriz[i])-1 {
		return buscaRecursiva(matriz, palavra, i, j+1)
	} else {
		return buscaRecursiva(matriz, palavra, i+1, 0)
	}
}

func BuscaMatriz() {
	matriz := [][]string{
		{"gato", "cachorro", "pássaro"},
		{"peixe", "coelho", "cavalo"},
		{"leão", "tigre", "elefante"},
	}

	palavra := "tigre"

	encontrado := buscaRecursiva(matriz, palavra, 0, 0)
	fmt.Printf("A palavra '%s' foi encontrada? %v\n", palavra, encontrado)
}
