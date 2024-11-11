package exercicios_revisao

import (
	"fmt"
)

type Livro struct {
	ID         int
	Titulo     string
	Autor      string
	Ano        int
	Emprestado bool
}

type Usuario struct {
	ID   int
	Nome string
}

type Emprestimo struct {
	UsuarioID int
	LivroID   int
}

var livros []Livro
var usuarios []Usuario
var historicoEmprestimos []Emprestimo

func cadastrarLivro() {
	var titulo, autor string
	var ano, id int
	fmt.Println("Digite o ID do livro:")
	fmt.Scanln(&id)
	fmt.Println("Digite o título do livro:")
	fmt.Scanln(&titulo)
	fmt.Println("Digite o autor do livro:")
	fmt.Scanln(&autor)
	fmt.Println("Digite o ano de publicação do livro:")
	fmt.Scanln(&ano)

	livro := Livro{
		ID:         id,
		Titulo:     titulo,
		Autor:      autor,
		Ano:        ano,
		Emprestado: false,
	}

	livros = append(livros, livro)
	fmt.Println("Livro cadastrado com sucesso!")
}

func cadastrarUsuario() {
	var id int
	var nome string
	fmt.Println("Digite o ID do usuário:")
	fmt.Scanln(&id)
	fmt.Println("Digite o nome do usuário:")
	fmt.Scanln(&nome)

	usuario := Usuario{
		ID:   id,
		Nome: nome,
	}

	usuarios = append(usuarios, usuario)
	fmt.Println("Usuário cadastrado com sucesso!")
}

func realizarEmprestimo() {
	var livroID, usuarioID int
	fmt.Println("Digite o ID do livro para empréstimo:")
	fmt.Scanln(&livroID)
	fmt.Println("Digite o ID do usuário:")
	fmt.Scanln(&usuarioID)

	var livroIndex = -1
	for i, livro := range livros {
		if livro.ID == livroID {
			livroIndex = i
			break
		}
	}

	if livroIndex == -1 || livros[livroIndex].Emprestado {
		fmt.Println("Livro não encontrado ou já emprestado!")
		return
	}

	var usuarioIndex = -1
	for i, usuario := range usuarios {
		if usuario.ID == usuarioID {
			usuarioIndex = i
			break
		}
	}

	if usuarioIndex == -1 {
		fmt.Println("Usuário não encontrado!")
		return
	}

	livros[livroIndex].Emprestado = true
	historicoEmprestimos = append(historicoEmprestimos, Emprestimo{
		UsuarioID: usuarioID,
		LivroID:   livroID,
	})

	fmt.Println("Empréstimo realizado com sucesso!")
}

func realizarDevolucao() {
	var livroID, usuarioID int
	fmt.Println("Digite o ID do livro para devolução:")
	fmt.Scanln(&livroID)
	fmt.Println("Digite o ID do usuário:")
	fmt.Scanln(&usuarioID)

	var livroIndex = -1
	for i, livro := range livros {
		if livro.ID == livroID {
			livroIndex = i
			break
		}
	}

	if livroIndex == -1 || !livros[livroIndex].Emprestado {
		fmt.Println("Livro não encontrado ou não foi emprestado!")
		return
	}

	var emprestimoIndex = -1
	for i, emprestimo := range historicoEmprestimos {
		if emprestimo.LivroID == livroID && emprestimo.UsuarioID == usuarioID {
			emprestimoIndex = i
			break
		}
	}

	if emprestimoIndex == -1 {
		fmt.Println("Este usuário não realizou o empréstimo desse livro!")
		return
	}

	livros[livroIndex].Emprestado = false
	historicoEmprestimos = append(historicoEmprestimos[:emprestimoIndex], historicoEmprestimos[emprestimoIndex+1:]...)
	fmt.Println("Devolução realizada com sucesso!")
}

func mostrarHistorico() {
	var livroID int
	fmt.Println("Digite o ID do livro para ver o histórico de empréstimos:")
	fmt.Scanln(&livroID)

	var livroIndex = -1
	for i, livro := range livros {
		if livro.ID == livroID {
			livroIndex = i
			break
		}
	}

	if livroIndex == -1 {
		fmt.Println("Livro não encontrado!")
		return
	}

	fmt.Println("Histórico de empréstimos do livro:", livros[livroIndex].Titulo)
	for _, emprestimo := range historicoEmprestimos {
		if emprestimo.LivroID == livroID {
			for _, usuario := range usuarios {
				if usuario.ID == emprestimo.UsuarioID {
					fmt.Printf("Usuário: %s\n", usuario.Nome)
				}
			}
		}
	}
}

func BibiotecaGO() {
	for {
		fmt.Println("\nMenu de Controle de Biblioteca:")
		fmt.Println("1. Cadastrar Livro")
		fmt.Println("2. Cadastrar Usuário")
		fmt.Println("3. Realizar Empréstimo")
		fmt.Println("4. Realizar Devolução")
		fmt.Println("5. Mostrar Histórico de Empréstimos de um Livro")
		fmt.Println("6. Sair")
		fmt.Print("Escolha uma opção: ")

		var escolha int
		fmt.Scanln(&escolha)

		switch escolha {
		case 1:
			cadastrarLivro()
		case 2:
			cadastrarUsuario()
		case 3:
			realizarEmprestimo()
		case 4:
			realizarDevolucao()
		case 5:
			mostrarHistorico()
		case 6:
			fmt.Println("Saindo do sistema...")
			return
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
