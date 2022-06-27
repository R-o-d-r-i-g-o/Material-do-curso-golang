package main

import "fmt"

// stack Overflow --> termo também utilizada pra se referir ao estouro de pilha
// Se uma função não tiver uma especificação de parada, ocorrerá o chamado "estouro de pilha" que armazenará cada
// vez mais dados no computador, provenientes dessas funções armazenadas (em formato de pilha) e acabará que o sistema não coseguirá comportar ela toda
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8 13
	posicao := uint(15)
	fmt.Println(fibonacci(posicao))

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

	// recursividade não é muito utilizada ultimamente por conta das estruturas de repetição
	// Não usar a recursão para tratativa de situações que muitas requisições sejam feitas
}
