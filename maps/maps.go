package maps

import "fmt"

func Maps() {
	//var m map[string]int
	//m = make(map[string]int)
	//
	//m = map[string]int{"samuel": 26, "gilberto": 42}
	//fmt.Println(m)

	//var pessoas map[int]string
	//pessoas = make(map[int]string)
	//pessoas = map[int]string{1: "samuel", 2: "gilberto"}
	//fmt.Println(pessoas)
	//
	//pessoas[1] = "Vanilton"
	//fmt.Println(pessoas)
	//
	//valor, exists := pessoas[1]
	//if exists {
	//	fmt.Println("valor existe e é", valor)
	//} else {
	//	fmt.Println("valor nao existente")
	//}
	//
	//for k, v := range pessoas {
	//	fmt.Println(k, v)
	//}
	//
	//delete(pessoas, 1)
	//fmt.Println(pessoas)
	//
	//pessoas[1] = "samuel"
	//fmt.Println(pessoas)
	//pessoas[3] = "ozzy"
	//fmt.Println(pessoas)

	var pessoas map[int]string
	pessoas = make(map[int]string)

	var id int
	var nome string

	for {
		fmt.Print("Digite o ID (ou 0 para sair): ")
		fmt.Scan(&id)

		if id == 0 {
			break
		}

		fmt.Print("Digite o nome: ")
		fmt.Scan(&nome)

		pessoas[id] = nome
	}
	for key, value := range pessoas {
		fmt.Println("Mapa final de pessoas", key, value)
	}

}
