package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	var x int
	fmt.Print("apresente o valor de x: ")
	fmt.Scan(&x)

	pertence := false
	for _, num := range lista {
		if num == x {
			pertence = true
			break
		}
	}
	if pertence {
		fmt.Print("x ja pertence à lista")
		fmt.Print(lista)
	} else {
		lista = append(lista, x)
		fmt.Print(lista)
	}

}
