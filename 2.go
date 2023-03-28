package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	x := 2

	lista = append(lista[:x], lista[x+1:]...)

	fmt.Print(lista)
}
