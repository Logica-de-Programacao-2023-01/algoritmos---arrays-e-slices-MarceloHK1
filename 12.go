package main

import "fmt"

func main() {
	lista := [5]int{3, 4, 6, 7, 9}
	var novalista []int

	for _, num := range lista {
		if num%3 == 0 {
			novalista = append(novalista, num)
		}
	}
	fmt.Print(novalista)
}
