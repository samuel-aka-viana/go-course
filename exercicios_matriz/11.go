package exercicios_matriz

import (
	"fmt"
	"math/rand"
	"time"
)

const tamanhoMatriz = 5

func imprimirMatriz(matriz [tamanhoMatriz][tamanhoMatriz]string) {
	for i := 0; i < tamanhoMatriz; i++ {
		for j := 0; j < tamanhoMatriz; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}

func verificarAtaque(matriz [tamanhoMatriz][tamanhoMatriz]string, x, y int) (bool, string) {
	if matriz[x][y] == "Z" {
		return true, "Você matou o zumbi!"
	}
	if matriz[x][y] == "J" {
		return true, "Você foi atacado pelo zumbi e morreu!"
	}
	return false, ""
}

func atacarJogador(matriz [tamanhoMatriz][tamanhoMatriz]string) {
	var linha, coluna int
	fmt.Println("Digite a linha (0-4) e a coluna (0-4) para atacar o zumbi:")
	fmt.Scanf("%d %d", &linha, &coluna)

	if linha < 0 || linha >= tamanhoMatriz || coluna < 0 || coluna >= tamanhoMatriz {
		fmt.Println("Coordenadas inválidas! Tente novamente.")
		return
	}

	atingiu, mensagem := verificarAtaque(matriz, linha, coluna)
	if atingiu {
		fmt.Println(mensagem)
	} else {
		fmt.Println("Você não acertou o zumbi. O jogo continua.")
	}

	matriz[linha][coluna] = "X"
}

func atacarZumbi(matriz *[tamanhoMatriz][tamanhoMatriz]string) {
	rand.Seed(time.Now().UnixNano())
	linha := rand.Intn(tamanhoMatriz)
	coluna := rand.Intn(tamanhoMatriz)

	matriz[linha][coluna] = "O"
	fmt.Printf("O zumbi atacou a posição (%d, %d)\n", linha, coluna)
}

func inicializarMatriz() [tamanhoMatriz][tamanhoMatriz]string {
	var matriz [tamanhoMatriz][tamanhoMatriz]string

	for i := 0; i < tamanhoMatriz; i++ {
		for j := 0; j < tamanhoMatriz; j++ {
			matriz[i][j] = "."
		}
	}

	matriz[0][0] = "J"
	matriz[4][4] = "Z"

	return matriz
}

func matrizZumbi() {
	matriz := inicializarMatriz()

	imprimirMatriz(matriz)

	jogadas := 0

	for {
		atacarJogador(matriz)
		imprimirMatriz(matriz)

		if matriz[0][0] == "O" {
			fmt.Println("Você foi atacado pelo zumbi e morreu!")
			break
		}

		atacarZumbi(&matriz)
		jogadas++
		imprimirMatriz(matriz)

		if matriz[4][4] == "O" {
			fmt.Println("O zumbi foi atacado e morreu! Você ganhou!")
			break
		}
	}

	fmt.Printf("O jogo terminou em %d jogadas!\n", jogadas)
}
