package main

import "fmt"

func soma(numeros ...int) int { // --> Opera de maneira a receber inumeros "numeros", neste caso, como parametro

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Não se pode haver mais de um paramentro variático por função
// Além disso, esse parametro variático deve ser o ultimo
func escrever(texto string, numeros ...int) {

	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	// totalDaSoma := soma() --> Se a função fosse passada vazia, o sitema entenderia um slice vazio e são tem problema algum
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}
