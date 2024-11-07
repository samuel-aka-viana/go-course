package exercicios_recursividade

import (
	"fmt"
)

func buscar(slice []string, alvo string) bool {
	if len(slice) == 0 {
		return false
	}
	if slice[0] == alvo {
		return true
	}
	return buscar(slice[1:], alvo)
}

func SearchString() {
	fmt.Println(buscar([]string{"go", "java", "python"}, "java"))
	fmt.Println(buscar([]string{"go", "java", "python"}, "ruby"))
}
