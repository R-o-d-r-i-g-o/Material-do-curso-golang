package main

import "fmt"

// go NÃO é uma lingaguem orientada a objetos e structs é forma que mais se aproxima de uma classe
// não existe de forma nativa herança, pelimorfismo, mas detêm extruturas que auxiliam nisso
type endereco struct {
	logradouro string
	numero     byte
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
	// por não terem recebido valor, já foram iniciadas automáticamente
	// com o chamado "valor zero"
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	// Maneira mais breve de realiar a mesma atribuição de valores dentro do struct
	usuario2 := usuario{"Davi", 22, enderecoExemplo}
	fmt.Println(usuario2)

	// maneira de passar apenas um valor como parametro
	// o mesmo pode ocorrer com a variavel idade, neste caso
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)

}
