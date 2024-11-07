package exercicios_recursividade

import (
	"fmt"
)

func invertido(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + invertido(s[:len(s)-1])
}

func StringInverted() {
	fmt.Println(invertido("golang"))
}
