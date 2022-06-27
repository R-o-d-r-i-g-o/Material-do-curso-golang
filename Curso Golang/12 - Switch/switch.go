package main

import "fmt"

func diasDaSemana(numero int) string {
	diaDaSemana := ""
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough // faz com que a condição abaixo seja ativada mesmo se não foi true
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	case 4:
		diaDaSemana = "Quarta"
	case 5:
		diaDaSemana = "Quinta"
	case 6:
		diaDaSemana = "Sexta"
	case 7:
		diaDaSemana = "Sábado"

	default:
		diaDaSemana = "Esta opção não existe"
	}

	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"

	default:
		return "Esta opção não existe"
	}
}

func main() {
	fmt.Println("Switch")

	// No switch case do golang, não há a necessidade de colocar break.
	if numero := 1; diasDaSemana(numero) == diaDaSemana2(numero) {
		fmt.Println(" O dia da semana é " + diasDaSemana(numero))
	} else {
		fmt.Println(" O dias da semanas são diferentes ")
	}

}
