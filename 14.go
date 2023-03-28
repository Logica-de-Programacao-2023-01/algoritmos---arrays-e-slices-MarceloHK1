package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x, y int
	fmt.Print("digite a posiçao dos termos que trocarao de lugar: ")
	fmt.Scan(&x, &y)

	lista[x], lista[y] = lista[y], lista[x]

	fmt.Print(lista)
}
