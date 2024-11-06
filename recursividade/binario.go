package recursividade

import "fmt"

func toBinary(n int) string {
	if n == 0 {
		return ""
	}
	return toBinary(n/2) + fmt.Sprintf("%d", n%2)
}

func Binario() {
	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	if numero == 0 {
		fmt.Println("Binário: 0")
	} else {
		binario := toBinary(numero)
		fmt.Printf("Binário: %s\n", binario)
	}
}
