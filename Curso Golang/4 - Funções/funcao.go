package main

import "fmt"

// toda a função que retornar algo deve ser tipada
// quando um valor for recebido como parametro não há necessidade de tipar ambos, apenas o últime se ambos forem semelhantes
func somar(n1, n2 byte) byte {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//calculoSoma, calculosSubtracao := calculosMatematicos(30, 20) // quando houver dois ou mais returns, deve haver o valor equivalente em funções que as recebem
	//fmt.Println(calculosSubtracao, calculoSoma)

	// pode se usar o _ para que não necessariamente alguma variável tenha que receber de dados vindos dos demais returns da função
	//calculoSoma, _ := calculosMatematicos(30, 20)
	//fmt.Println(calculoSoma)

	_, calculoSubtracao := calculosMatematicos(30, 20)
	fmt.Println(calculoSubtracao)
}
