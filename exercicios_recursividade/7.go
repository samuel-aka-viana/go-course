package exercicios_recursividade

import "fmt"

func somaValores(mapa map[string]interface{}) int {
	total := 0

	for _, valor := range mapa {
		switch v := valor.(type) {
		case int:
			total += v
		case map[string]interface{}:
			total += somaValores(v)
		}
	}

	return total
}

func NestedMap() {
	mapa := map[string]interface{}{
		"a": 2,
		"b": map[string]interface{}{
			"c": 3,
			"d": 4,
		},
	}

	fmt.Println(somaValores(mapa)) // Output: 9
}
