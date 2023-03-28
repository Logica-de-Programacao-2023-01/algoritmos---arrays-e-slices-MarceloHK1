package main

import "fmt"

func main() {
	matriz := [3][2]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print("digite o valor da linha ", i+1, "e o valor da coluna ", j+1, ":")
			fmt.Scan(&matriz[i][j])
		}
	}
	fmt.Print(matriz)

}
