package main

import "fmt"

func main() {
	lista := [7]int{1, 2, 3, 4, 5, 6, 7}
	var x int
	fmt.Print("digite o valor de x que sera adicionado no primeiro e ultimo termo da lista: ")
	fmt.Scan(&x)

	lista[0] += x
	lista[len(lista)-1] += x

	fmt.Print(lista)
}
