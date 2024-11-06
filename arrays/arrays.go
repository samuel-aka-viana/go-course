package arrays

import "fmt"

func Arrays() {
	var arr [5]int // inteiro
	fmt.Println(arr)

	//var array_flora [4]float64{1.3,2.}
	//fmt.Println(array_flora)

	var arr_inteiros = [3]int{1, 2, 3}
	fmt.Println(arr_inteiros)

	var arr_string [4]string
	fmt.Println(arr_string)

	var arr_linguagens = [4]string{"Go", "Python", "Java", "C++"}
	fmt.Println(arr_linguagens)

	// mudanças no array

	arr[0] = 10
	fmt.Println(arr)

	// iterar arrays

	for i := 0; i < len(arr_inteiros); i++ {
		fmt.Println(arr_inteiros[i])
	}
	//
	for index, value := range arr_inteiros {
		fmt.Println(index, value)
	}

	var frutas = [4]string{"melancia", "laranja", "limao", ""}
	fmt.Println(frutas)

	for _, value := range frutas {
		fmt.Println(value)
	}
	//
}
