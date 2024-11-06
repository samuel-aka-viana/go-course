package exercicio_slice

import "fmt"

func MenuGo() {
	var nomes []string
	var sexos []string
	var opcao int

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Adicionar nome e sexo de uma pessoa")
		fmt.Println("2. Listar todos os nomes")
		fmt.Println("3. Total de homens e total de mulheres")
		fmt.Println("4. Sair do programa")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 1 {
			var nome, sexo string
			fmt.Print("Digite o nome: ")
			fmt.Scan(&nome)
			fmt.Print("Digite o sexo (Masculino/Feminino): ")
			fmt.Scan(&sexo)
			nomes = append(nomes, nome)
			sexos = append(sexos, sexo)
			fmt.Println("Pessoa adicionada com sucesso!")

		} else if opcao == 2 {
			fmt.Println("\nLista de pessoas:")
			for i := 0; i < len(nomes); i++ {
				fmt.Printf("Nome: %s, Sexo: %s\n", nomes[i], sexos[i])
			}

		} else if opcao == 3 {
			totalHomens, totalMulheres := 0, 0
			for i := 0; i < len(sexos); i++ {
				if sexos[i] == "Masculino" {
					totalHomens++
				} else if sexos[i] == "Feminino" {
					totalMulheres++
				}
			}
			fmt.Printf("\nTotal de homens: %d\n", totalHomens)
			fmt.Printf("Total de mulheres: %d\n", totalMulheres)

		} else if opcao == 4 {
			fmt.Println("Saindo do programa...")
			break

		} else {
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
