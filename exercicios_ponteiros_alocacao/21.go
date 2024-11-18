package exercicios_ponteiros_alocacao

import "fmt"

func atualizarMapa(mapa *map[string]int, chave string) {
	(*mapa)[chave] += 10
}

func UpdateMap() {
	m := map[string]int{
		"a": 5,
		"b": 15,
		"c": 20,
	}

	fmt.Println("Antes:", m)

	atualizarMapa(&m, "a")

	fmt.Println("Depois:", m)
}
