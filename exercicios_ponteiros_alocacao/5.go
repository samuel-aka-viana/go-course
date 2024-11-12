package exercicios_ponteiros_alocacao

import "fmt"

func alocarSlice(n int) *[]int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return &slice
}

func AllocSlice() {
	slice1 := alocarSlice(5)
	slice2 := alocarSlice(8)
	slice3 := alocarSlice(3)

	fmt.Println("Slice 1:", *slice1)
	fmt.Println("Slice 2:", *slice2)
	fmt.Println("Slice 3:", *slice3)
}
