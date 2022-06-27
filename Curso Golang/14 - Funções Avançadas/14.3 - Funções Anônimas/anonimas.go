package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10) // --> função que pode ser usada para concatenaa algumas informações
	}("Passando Parâmetro") // --> Esses parenteses devem sempre aparecer no final de uma função anônima

	fmt.Print(retorno)
}
