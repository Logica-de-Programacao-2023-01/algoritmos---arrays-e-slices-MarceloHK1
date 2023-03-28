package main

import "fmt"

func main() {
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0

	for i := 0; i < len(lista); i += 2 {
		soma += lista[i]
	}
	fmt.Print("a soma dos elementos em posiçoes pares é ", soma)
}
