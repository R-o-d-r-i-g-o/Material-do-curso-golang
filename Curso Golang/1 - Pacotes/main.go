package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("rodrigomarqribeiro@gmail.com")
	fmt.Println(erro)
}

// rodar no console: "go run main.go"
// para atualizar e rodar a principio usa-se o comando: "go build"
// "go install" direciona o arquivo go.mod para onde foi instalado o golang
// go get github.com/badoux/checkmail - importou um pacote externo para a validação de e-mails
// go mod tidy - retira todas as dependencias em desuso
