package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")
	numero := 10

	// Não existe operador ternário em go, logo as estruturas abaixo devem ser mantidas
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// Por meio dessa estrutura abaixo, é possível inicializar uma variável direto no if
	// seu único porém é que se restringe somente ao escopo do if e outras estruturas encadiados, se ouver.
	// estrutura chamada "if init"
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que zero")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}
