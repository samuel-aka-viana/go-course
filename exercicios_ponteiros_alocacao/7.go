package exercicios_ponteiros_alocacao

import "fmt"

func adicionarRegistro(m *map[string]int, chave string, valor int) {
	(*m)[chave] = valor
}

func AddRegister() {
	registros := make(map[string]int)

	adicionarRegistro(&registros, "Alice", 25)
	adicionarRegistro(&registros, "Bob", 30)
	adicionarRegistro(&registros, "Carol", 27)

	fmt.Println("Mapa atualizado:", registros)
}
