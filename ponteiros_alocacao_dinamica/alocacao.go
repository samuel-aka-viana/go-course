package ponteiros_alocacao_dinamica

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Alloc() {
	person1 := new(*Person)
	fmt.Println(&person1)

	myMap := make(map[string]int)
	myMap["chave"] = 26
	fmt.Println(myMap)

	inteiro := new(int)
	fmt.Println(*inteiro)

	numbers := []int{10, 20, 30, 40, 50}
	p := &numbers

	first := (*p)[0]
	fmt.Println("First element:", first)

	second := (*p)[1]
	fmt.Println("Second element:", second)

	third := (*p)[2]
	fmt.Println("Third element:", third)

	for i := 0; i < len(*p); i++ {
		fmt.Printf("Element %d: %d\n", i, (*p)[i])
	}

}
