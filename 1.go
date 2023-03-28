package main

import "fmt"

func main() {
	lista := [3]int{1, 2, 3}
	soma := 0

	for _, list := range lista {
		soma += list
	}

	fmt.Print("a soma do array é ", soma)
}
