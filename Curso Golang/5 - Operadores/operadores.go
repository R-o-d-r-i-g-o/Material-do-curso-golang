package main

import "fmt"

func main() {
	// Aritiméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Não se pode realizar nenhuma opração com variaveis tipadas com diferentes formas.
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	// Fim dos aritiméticos

	// Atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// Fim dos operadores de atribuicao

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos operadores relacionais

	// Operadores lógicos
	fmt.Println("------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// Fim dos operadores lógicos

	// Operadores unários
	// Em go esses operadores unários só existem de maneira pós fixada
	// opradores de auto-incremento e decremento, etc.:
	numero := 10
	numero++     //numero = numero + 1
	numero += 15 //numero = numero + 15
	numero--     //numero = numero - 1
	numero -= 15 //numero = numero - 15
	fmt.Println(numero)

	numero *= 15 //numero = numero * 15
	numero /= 15 //numero = numero / 15
	numero %= 15 //numero = numero % 15
	fmt.Println(numero)
	// Fim dos operadores unários

	// Estrutura de seleção
	// Em go não existe operador ternário, logo é-se necessário usar a extrutura abaixo:
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
	// Fim da estrutura de seleção

}
