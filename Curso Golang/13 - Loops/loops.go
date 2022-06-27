package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	// Em golang existe apenas uma estrutura de repetição (for), que pode ser moldada para se adequar à diversas possibilidades
	// tais como o if init.
	for i < 10 {
		i++
		fmt.Println(i)

		// Esta função é usada para diminuir velocidade do sistema, colocando
		// um intervalo de tempo para até próxima ocorrência de um evento.
		time.Sleep(time.Second)
	}

	// j := 0; j < 10; j += 5 --> também e possível
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	// Seguindo essa estrutura, o range percorrerá a string
	// esse _ representa o não recebimento do valor do indice, obrigatorio nessa estrutura
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Seguindo essa estrutura, o range percorrerá a string
	for indice, letra := range "PALAVRA" {
		// A estrutura "string(letra)" foi usada para retorna as elas e não códigos da tabela ascii.
		fmt.Println(indice, string(letra))
	}

	// Essa estrutura range funciona perfeitamente em métodos, mas não em structs
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// type usuarioStruct struct {
	// 	nome string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct{"Zé", "Júnior"}
	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }


	for {  // --> Executa o loop infinitamente
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
