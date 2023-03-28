package main

import "fmt"

func main() {
	lista := [6]float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.6}
	var x float64
	fmt.Print("digite o valor de x: ")
	fmt.Scan(&x)

	for i := 0; i < len(lista); i++ {
		lista[i] += x
	}

	fmt.Print(lista)

}
