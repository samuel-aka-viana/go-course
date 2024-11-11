package exercicios_revisao

import (
	"fmt"
)

type Item struct {
	ID         int
	Nome       string
	Preco      float64
	Quantidade int
}

type Vendedor struct {
	ID          int
	Nome        string
	TotalVendas float64
}

var estoque []Item
var vendedores map[int]Vendedor

func saldoTotalRecursivo(estoque []Item, index int, saldoTotal float64) float64 {
	if index >= len(estoque) {
		return saldoTotal
	}

	item := estoque[index]
	saldoTotal += item.Preco * float64(item.Quantidade)

	return saldoTotalRecursivo(estoque, index+1, saldoTotal)
}

func saldoVendedorRecursivo(vendedores map[int]Vendedor, vendedorID int, saldoVendedor float64) float64 {
	if vendedor, exists := vendedores[vendedorID]; exists {
		saldoVendedor += vendedor.TotalVendas
	}
	return saldoVendedor
}

func cadastrarItem() {
	var nome string
	var preco float64
	var quantidade int
	fmt.Println("Digite o nome do item:")
	fmt.Scanln(&nome)
	fmt.Println("Digite o preço do item:")
	fmt.Scanln(&preco)
	fmt.Println("Digite a quantidade do item:")
	fmt.Scanln(&quantidade)

	item := Item{
		ID:         len(estoque) + 1,
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}

	estoque = append(estoque, item)
	fmt.Println("Item cadastrado com sucesso!")
}

func removerItem() {
	var id int
	fmt.Println("Digite o ID do item a ser removido:")
	fmt.Scanln(&id)

	var index = -1
	for i, item := range estoque {
		if item.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Item não encontrado!")
		return
	}

	estoque = append(estoque[:index], estoque[index+1:]...)
	fmt.Println("Item removido com sucesso!")
}

func cadastrarVendedor() {
	var nome string
	fmt.Println("Digite o nome do vendedor:")
	fmt.Scanln(&nome)

	vendedor := Vendedor{
		ID:          len(vendedores) + 1,
		Nome:        nome,
		TotalVendas: 0,
	}

	vendedores[vendedor.ID] = vendedor
	fmt.Println("Vendedor cadastrado com sucesso!")
}

func realizarVenda() {
	var vendedorID, itemID, quantidade int
	fmt.Println("Digite o ID do vendedor:")
	fmt.Scanln(&vendedorID)
	fmt.Println("Digite o ID do item:")
	fmt.Scanln(&itemID)
	fmt.Println("Digite a quantidade vendida:")
	fmt.Scanln(&quantidade)

	var itemIndex = -1
	for i, item := range estoque {
		if item.ID == itemID {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 || estoque[itemIndex].Quantidade < quantidade {
		fmt.Println("Item não encontrado ou estoque insuficiente!")
		return
	}

	estoque[itemIndex].Quantidade -= quantidade

	vendedor, exists := vendedores[vendedorID]
	if !exists {
		fmt.Println("Vendedor não encontrado!")
		return
	}

	vendedor.TotalVendas += estoque[itemIndex].Preco * float64(quantidade)
	vendedores[vendedorID] = vendedor

	fmt.Println("Venda realizada com sucesso!")
}

func mostrarSaldoTotal() {
	total := saldoTotalRecursivo(estoque, 0, 0)
	fmt.Printf("Saldo total do estoque: R$ %.2f\n", total)
}

func mostrarSaldoVendedor() {
	var vendedorID int
	fmt.Println("Digite o ID do vendedor:")
	fmt.Scanln(&vendedorID)

	saldo := saldoVendedorRecursivo(vendedores, vendedorID, 0)
	fmt.Printf("Saldo do vendedor %d: R$ %.2f\n", vendedorID, saldo)
}

func mostrarEstoque() {
	if len(estoque) == 0 {
		fmt.Println("Estoque vazio!")
		return
	}

	fmt.Println("Itens em estoque:")
	for _, item := range estoque {
		fmt.Printf("ID: %d | Nome: %s | Preço: R$ %.2f | Quantidade: %d\n", item.ID, item.Nome, item.Preco, item.Quantidade)
	}
}

func StockControl() {
	estoque = make([]Item, 0)
	vendedores = make(map[int]Vendedor)

	for {
		fmt.Println("\nMenu de Controle de Estoque:")
		fmt.Println("1. Cadastrar Item")
		fmt.Println("2. Remover Item")
		fmt.Println("3. Cadastrar Vendedor")
		fmt.Println("4. Realizar Venda")
		fmt.Println("5. Mostrar Saldo Total do Estoque")
		fmt.Println("6. Mostrar Saldo de Vendedor")
		fmt.Println("7. Mostrar Itens em Estoque")
		fmt.Println("8. Sair")
		fmt.Print("Escolha uma opção: ")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			cadastrarItem()
		case 2:
			removerItem()
		case 3:
			cadastrarVendedor()
		case 4:
			realizarVenda()
		case 5:
			mostrarSaldoTotal()
		case 6:
			mostrarSaldoVendedor()
		case 7:
			mostrarEstoque()
		case 8:
			fmt.Println("Saindo do sistema...")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
