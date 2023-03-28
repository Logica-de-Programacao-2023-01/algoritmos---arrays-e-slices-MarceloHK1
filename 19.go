package main

import "fmt"

func main() {
	var x int
	fmt.Print("digite o tamanho dos arrays: ")
	fmt.Scan(&x)

	lista1 := make([]int, x)
	fmt.Print("digite os valores da lista1: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&lista1[i])
	}
	lista2 := make([]int, x)
	fmt.Print("digite os valores da lista2: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&lista2[i])
	}
	lista3 := make([]int, x)
	for i := 0; i < x; i++ {
		lista3[i] = lista1[i] + lista2[i]
	}
	fmt.Println(lista3)
}
