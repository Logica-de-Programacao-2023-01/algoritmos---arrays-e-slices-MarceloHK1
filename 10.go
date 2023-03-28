package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min := lista[0]
	max := lista[0]

	for _, num := range lista {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println("o valor minimo é ", min)
	fmt.Println("o valor maximo é ", max)
}
