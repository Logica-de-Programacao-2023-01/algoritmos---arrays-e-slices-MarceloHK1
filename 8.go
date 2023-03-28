package main

import "fmt"

func main() {
	lista := []string{"mar", "ce", "lo", "hon", "da", "ko", "ba", "ya"}

	var remover string
	fmt.Print("insira o valor removivel ")
	fmt.Scan(&remover)

	for i := 0; i < len(lista); i++ {
		if lista[i] == remover {
			lista = append(lista[:i], lista[i+1:]...)
			i--
		}
	}
	fmt.Print("a lista após a remoção é ", lista)

}
