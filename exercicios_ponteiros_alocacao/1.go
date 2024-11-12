package exercicios_ponteiros_alocacao

import "fmt"

func doubleValue(value *int) {
	*value = *value * 2
}

func AlterReference() {
	variable := 10
	fmt.Printf("valor original %d", variable)
	doubleValue(&variable)
	fmt.Printf("valor dobrado %d\n", variable)

}
