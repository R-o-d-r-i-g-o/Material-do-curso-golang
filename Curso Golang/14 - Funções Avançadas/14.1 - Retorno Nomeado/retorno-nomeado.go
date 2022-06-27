package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return // --> por já ser expecificado o retorno, não há a necessidade de por os retornos aqui
}

// A maneira abaixo também é uma opção, contudo não tão prática
// func CalculosMatematicos(n1, n2 int) (int, int) {
// 	soma := n1 + n2
// 	subtracao := n1 - n2
// 	return soma, subrtracao
// }

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Print(soma, subtracao)
}
