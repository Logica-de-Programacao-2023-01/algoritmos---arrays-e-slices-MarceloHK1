package main

import "fmt"

func main() {
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var x int
	fmt.Print("digite um valor de x ")
	fmt.Scan(&x)

	pertence := false
	for _, y := range lista {
		if y == x {
			pertence = true
			break
		}
	}
	if pertence {
		fmt.Print("o numero", x, "esta presente na lista")
	} else {
		fmt.Print("o numero", x, "nao esta presente na lista")
	}
}
