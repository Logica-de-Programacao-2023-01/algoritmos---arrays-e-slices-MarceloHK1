package main

import "fmt"

func main() {
	var x int
	fmt.Print("digite um valor de x: ")
	fmt.Scan(&x)

	primos := make([]int, 0, x)
	i := 2
	for len(primos) < x {
		if primo(i) {
			primos = append(primos, i)
		}
		i++
	}
	fmt.Println(primos)
}

func primo(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
