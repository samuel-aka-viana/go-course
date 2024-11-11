package exercicios_revisao

import "fmt"

func invertMap(original map[string]int) map[int]string {
	invertedMap := make(map[int]string)

	for name, age := range original {
		invertedMap[age] = name
	}

	return invertedMap
}

func InvertedMap() {
	namesToAges := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"David":   28,
	}

	inverted := invertMap(namesToAges)

	fmt.Println("Map invertido (idade -> nome):")
	for age, name := range inverted {
		fmt.Printf("Idade: %d, Nome: %s\n", age, name)
	}
}
