package exercicios_revisao

import "fmt"

func removerChave(m map[string]int, chave string) map[string]int {
	delete(m, chave)
	return m
}

func RemoveMap() {
	meuMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	fmt.Println("Map Original:", meuMap)

	meuMap = removerChave(meuMap, "b")

	fmt.Println("Map após remover a chave 'b':", meuMap)
}
