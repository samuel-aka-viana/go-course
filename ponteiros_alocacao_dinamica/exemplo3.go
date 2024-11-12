package ponteiros_alocacao_dinamica

import "fmt"

func addOrRemoveKey(m *map[string]int, key string, value int) {
	(*m)[key] = value
}

func removeKey(m *map[string]int, key string) {
	delete(*m, key)
}

func AddRemoveMap() {
	myMap := map[string]int{"a": 1, "b": 2}
	fmt.Println("map original", myMap)
	addOrRemoveKey(&myMap, "c", 3)
	fmt.Println("adicionado o map", myMap)
	addOrRemoveKey(&myMap, "c", 4)
	fmt.Println("atualizando a key do map", myMap)
	removeKey(&myMap, "a")
	fmt.Println("removado o map", myMap)
}
