package exercicios_revisao

import "fmt"

func toHex(n int) string {
	if n == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"

	if n < 16 {
		return string(hexChars[n])
	}

	return toHex(n/16) + string(hexChars[n%16])
}

func Hex() {
	n := 255
	hexResult := toHex(n)

	fmt.Printf("O número %d em hexadecimal é: %s\n", n, hexResult)
}
