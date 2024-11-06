package exercicios_struct

import "fmt"

// Implemente um programa em Go para simular partidas
// de futebol entre times em um pequeno campeonato,
// onde cada time acumula saldo de gols ao longo de várias rodadas.
// No final, o programa deverá determinar o vencedor com base no saldo de gols.
type Time struct {
	Nome      string
	SaldoGols int
}

func jogarRodada(times []Time) {
	fmt.Println("\nJogos da rodada:")

	placares := []struct {
		time1Index int
		time2Index int
		golsTime1  int
		golsTime2  int
	}{
		{0, 1, 2, 1}, // Flamengo 2 x 1 Palmeiras
		{2, 3, 1, 3}, // Atlético-MG 1 x 3 Barcelona
		{4, 5, 2, 2}, // Real Madrid 2 x 2 Manchester City
	}

	for _, jogo := range placares {
		time1 := times[jogo.time1Index]
		time2 := times[jogo.time2Index]

		fmt.Printf("%s %d x %d %s\n", time1.Nome, jogo.golsTime1, jogo.golsTime2, time2.Nome)

		time1.SaldoGols += (jogo.golsTime1 - jogo.golsTime2)
		time2.SaldoGols += (jogo.golsTime2 - jogo.golsTime1)

		times[jogo.time1Index] = time1
		times[jogo.time2Index] = time2
	}
}

func determinarVencedor(times []Time) Time {
	vencedor := times[0]
	for _, time := range times {
		if time.SaldoGols > vencedor.SaldoGols {
			vencedor = time
		}
	}
	return vencedor
}

func DesafioVencedor() {
	times := []Time{
		{Nome: "Flamengo", SaldoGols: 0},
		{Nome: "Palmeiras", SaldoGols: 0},
		{Nome: "Atlético-MG", SaldoGols: 0},
		{Nome: "Barcelona", SaldoGols: 0},
		{Nome: "Real Madrid", SaldoGols: 0},
		{Nome: "Manchester City", SaldoGols: 0},
	}

	rodadas := 4
	for rodada := 1; rodada <= rodadas; rodada++ {
		fmt.Printf("\nRodada %d\n", rodada)
		jogarRodada(times)
	}

	fmt.Println("\nSaldo de gols final de cada time:")
	for _, time := range times {
		fmt.Printf("Time: %s, Saldo de Gols: %d\n", time.Nome, time.SaldoGols)
	}

	vencedor := determinarVencedor(times)
	fmt.Printf("\nO vencedor é %s com saldo de gols de %d!\n", vencedor.Nome, vencedor.SaldoGols)
}
