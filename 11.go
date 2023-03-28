package main

import "fmt"

func main() {
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var i, j int
	fmt.Print("digite a linha  e a coluna ")
	fmt.Scan(&i, &j)

	fmt.Printf("o numero na linha (%d) e na coluna (%d) é %d\n ", i, j, matriz[i][j])
}
