package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")
	// Diferença entre array e slice:
	// Em go ambos simbolizam a mesma coisa, uma espécie de trem com diversos vagões
	// de semelhante tipagem que armazenam coisas. Contudo, array apresenta um número
	// de vagões fixo, de maneira geral, pré-determinado. Já os slices uma quantidade dinâmica

	// Outra diferença técnica é que o slice referência um array, ou seja, ele não é uma opção
	// mediante ao uso do array ou não, pode-se dizer que ele é um sub-tipo do mesmo.

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// ele fixará a quantidade de itens a medida que forem acressentados
	// Obs.: essa notação não deixa o array com tamanho dinâmico
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// também chamado array interno.
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// O é um pacote reflect que retorna o tipo da variável, e em caso que são envolvidos
	// array e slices, ele retorna também o número de vagões (indices) que esses tem, anda na dinamica do trem
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// A função append() pode ser usada somente em slices
	slice = append(slice, 18)
	fmt.Println(slice)

	// esta é uma maneira de selecionar intervalos dentro de um array e / ou slice
	// o valor inicial delimita o início do intervalo e o final seu término
	slice2 := array2[1:3]
	// Ainda no caso acima, a variável "slice2" apenas se tornou um slice, pois, como o próprio nome sugere,
	// representa um pedaço, neste caso de array. Logo ele está "apontando" para um intervalo especifico do array.
	slice2 = append(slice2, "gfdgdgfd")
	fmt.Println(reflect.TypeOf(slice2))
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Arrays Internos
	// a função make gera, a principio, um array com todos os indices recebendo valor zero.
	// Depois de acrescentados novos valores, que ocuparão outros indices, ele se tornará automaticamente um slice.
	// como parâmetro, primeiro é passado o tipo, depois o número mínimo de indices permiditidos
	// e logo após o máximo, ou seja, até onde ele pode se extender futuramente.
	slice3 := make([]float32, 10, 15)
	fmt.Println(reflect.TypeOf(slice3))
	fmt.Println(slice3)

	// a função len() conta o número de indices do array / slice, semelhante á func. length em outras linguagens
	fmt.Println(len(slice3)) // length

	// a função cap() retorna a capacidade máxima permitida pelo slice "slice3"
	fmt.Println(cap(slice3)) // capacidade

	// Dessa maneira, a função make apenas delimita um index mínimo.
	// Mas, em geral, não é necessário determinar um indice máximo, pois a medida que ele é
	// ultrapassado ele aumenta o valor mínimo para a respectiva quantia e duplica-se, tendo como base
	// o número mínimo de indices, e o acrescenta no local do valor máximo. Dessa forma nunca tendo um limite
	slice4 := make([]float32, 5)
	fmt.Println(slice4)

	// aqui é onde ocorre tudo aquilo dito no comentário acima
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	// *Obs.: Sempre que um slice é criado, ele referencia a fatia de um array.
	// Ressumo: Array é uma lista de tamanho fixo ao contrário do Slice.
}
