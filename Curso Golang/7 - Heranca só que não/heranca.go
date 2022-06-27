package main

import (
	"fmt"
)

// o go permite que sejam criados tipos que não serão usados
// por se assimilarem aos "tipos" de variaveis
func main() {
	fmt.Println("Herança")

	//var p1 pessoa
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	// var e1 estudante
	// Sempre deve se chamar a struct que simboliza a herança na que a herdou
	e1 := estudante{"Engenharia", "USP", p1}
	fmt.Println(e1)

	// também pode-se chamar dessa maneira um valor especifico
	// caso dentro de estudante a struct pessoa fosse tipada, deveria usar o
	// seguinte atalho e1.pessoa.nome, seguindo dessa maneira -> fmt.Println(e1.pessoa.nome)
	fmt.Println(e1.nome)

}

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// por não ser desejado que o caminho estudante.pessoa.nome, por exemplo seja feito
// foi optado colocar apenas a variável pessoa, sem seu respectivo tipo "pessoa" na função estudante
type estudante struct {
	curso     string
	faculdade string
	pessoa
}
