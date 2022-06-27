package main

import "fmt"

// Ponteiro não salva necessariamente um valor, mas um local na memória
//  que detêm esse valor. Logo, ele aponta o valor
func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	// O valor zero de um ponteiro é NULL, em go <nil>
	// Usa-se * antes do tipo da variável para se identificar um ponteiro
	// depois para o ponteiro receber uma variável ela deve ser precedida de &

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro)
	// Se apenas aparente o nome da variável, chamada ponteiro, ele apena irá
	// localizar o valor na memória com já de costume, contudo se haver um * precedendo
	// a variável, ele receberá o real valor da mesma.
	// Esse * antes da váriavel é chamado de "Desreferenciação"

	variavel3 = 150
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)

}
