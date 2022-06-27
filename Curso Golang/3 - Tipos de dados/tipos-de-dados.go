package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64 - tipos de tipagem de números inteiros
	// quando for digitado apenas "int" na tipagem, este será correspondente ao valor da própria máquina
	// ex.: windowns 32 terá int32 automaticamente
	var numero int16 = -100
	fmt.Println(numero)

	//uint8, uint16, uint32, uint64 - tipos de tipagem de números inteiros sem sinal
	//unit -> unsygned int
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//int32 = rune -> Contudo, geralmente usado para representar números da table ascii
	//
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = unit8
	var numero4 byte = 123
	fmt.Println(numero4)

	// no caso dos numeros com ponto flutuante, não podem ser tipados somente com float
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1200000003.450000000000
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.67
	fmt.Println(numeroReal3)

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	// se um caracter é declarada dentro de aspas simples ele logo é tranformado em numero frente a tabela ascii
	// não existe o char em golang
	char := 'p'
	fmt.Println(char)

	// desta maneira qualquer variável já é declarada com o valor zero independentemente de sua aparição
	var texto int16
	fmt.Println(texto)

	// desta maneira, obrigamente a variável precisa ser iniciada
	texto1 := 0
	fmt.Println(texto1)

	// o valor zero de uma variável boolean é false
	var booleano1 bool
	fmt.Println(booleano1)

	// valor nulo
	var erro error
	fmt.Println(erro) // <nil>, no caso do golang

	// caso necessário atribuir algum valor à uma variável error, usa-se o pacote interno "errors"
	var erro1 error = errors.New("Erro Interno")
	fmt.Println(erro1)
}
