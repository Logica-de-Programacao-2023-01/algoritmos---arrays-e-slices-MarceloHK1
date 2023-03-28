package main

import "fmt"

func main() {
	var x int
	fmt.Print("digite o tamanho do array: ")
	fmt.Scan(&x)

	lista := make([]int, x)
	fmt.Print("digite os valores do array: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&lista[i])
	}
	teste := true
	for i := 0; i < x-1; i++ {
		if lista[i] > lista[i+1] {
			teste = false
			break
		}
	}
	if teste {
		fmt.Print("a lista esta em ordem crescente")
	} else {
		fmt.Print("a lista nao esta em ordem crescente")
	}
}
