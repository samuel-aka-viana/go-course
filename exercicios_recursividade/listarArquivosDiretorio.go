package exercicios_recursividade

import (
	"fmt"
	"os"
	"path/filepath"
)

func listarArquivos(diretorio string) {
	dir, err := os.Open(diretorio)
	if err != nil {
		fmt.Println("Erro ao abrir diretório:", err)
		return
	}
	defer dir.Close()

	arquivos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
		return
	}

	for _, arquivo := range arquivos {
		caminho := filepath.Join(diretorio, arquivo.Name())

		if arquivo.IsDir() {
			fmt.Println("Diretório:", caminho)
			listarArquivos(caminho)
		} else {
			fmt.Println("Arquivo:", caminho)
		}
	}
}

func ShowFiles() {
	listarArquivos("./")
}
