package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Print("qual o tamanho do slice? ")
	fmt.Scan(&size)

	lista := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Print("digite o elemento", i+1, ": ")
		fmt.Scan(&lista[i])
	}
	fmt.Print("lista: ", lista)

	soma := 0
	for _, y := range lista {
		soma += y
	}
	fmt.Print("a soma dos elementos do slice é: ", soma)
}
