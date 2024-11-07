package exercicios_recursividade

import "fmt"

func adivinharNumero(array []int, indice int, palpite int, tentativa int) {
	if array[indice] == palpite {
		fmt.Println("Parabéns! Você acertou o número.")
		return
	}

	if tentativa < indice {
		fmt.Println("Você está muito longe. Tente um número maior!")
	} else if tentativa > indice {
		fmt.Println("Você está muito longe. Tente um número menor!")
	} else {
		fmt.Println("Você está perto!")
	}

	var novoPalpite int
	fmt.Print("Tente outro número: ")
	fmt.Scan(novoPalpite)
	adivinharNumero(array, indice, novoPalpite, tentativa)
}

func GameRecursion() {
	var n int
	fmt.Print("Quantos números você quer no array? ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Preencha o array com números:")
	for i := 0; i < n; i++ {
		fmt.Printf("Digite o número %d: ", i+1)
		fmt.Scan(array[i])
	}

	var indice int
	fmt.Print("Escolha o índice de onde está o número que a outra pessoa vai tentar adivinhar: ")
	fmt.Scan(indice)

	if indice < 0 || indice >= n {
		fmt.Println("Índice inválido!")
		return
	}

	var palpite int
	fmt.Print("Tente adivinhar o número: ")
	fmt.Scan(palpite)

	adivinharNumero(array, indice, palpite, indice)
}
