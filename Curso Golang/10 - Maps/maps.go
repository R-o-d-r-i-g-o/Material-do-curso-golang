package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Map é uma strutura bem similar ao struct
	// Mais similar ainda a objeto e / ou vetor associativo
	usuario1 := map[int]string{
		1: "Pedro",
		2: "Silva",
	}

	usuario2 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus",
		},
	}

	fmt.Println(usuario1[1])      // Maneira espefica de captar um valor
	fmt.Println(usuario2["nome"]) // Maneira espefica de captar um valor
	fmt.Println(usuario2)

	// Esses colchetes são similares aos pontos no js
	fmt.Println(usuario3["curso"]["primeiro"])
	fmt.Println(usuario3["curso"])

	// função que deleta as chaves de um map
	delete(usuario2, "nome")
	fmt.Println(usuario3)

	// Atena-se à estrutura original do map para adinionar novas chaves.
	// Também é possível adicionar mais chaves no map pela seguinte estrutura:
	usuario3["signo"] = map[string]string{
		"nome": "Gêmenos",
	}
	fmt.Println(usuario3)
}
