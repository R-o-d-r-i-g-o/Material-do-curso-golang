package main

import "fmt"

// funções que começam com letra maiúscula podem ser importadas em outros arquivos
func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variável 2" // maneira de declarar uma variável, com tipagem automática
	// esse sinal := recebe o nome de inferência

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( // maneira de declarar duas ou mais variáveis de uma vez só
		variavel3 string = "lalaal"
		variavel4 string = "lalla"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel3", "variavel4"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
