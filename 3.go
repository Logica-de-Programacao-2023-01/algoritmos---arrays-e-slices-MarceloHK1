package main

import "fmt"

func main() {
	lista := [4]float64{1.2, 1.3, 1.4, 1.5}
	produto := 1.0

	for _, list := range lista {
		produto *= list
	}

	fmt.Print(produto)
}
