package main

import "fmt"

func main() {
	lista := [10]float64{1.0, 1.1, 1.3, 1.4, 1.5, 1.6, 5.3, 5.4, 5.0, 10.6}
	var maior []float64

	for _, num := range lista {
		if num > 5 {
			maior = append(maior, num)
		}
	}
	fmt.Print(maior)

}
