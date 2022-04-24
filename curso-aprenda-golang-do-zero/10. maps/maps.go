package main

import "fmt"

func main() {
	fmt.Println("maps")

	// o map é uma forma de struct muito mais rígido, pois todas as chaves serão do tipo A e todos os valores do tipo B
	// ao contrário do struct, o map não utiliza map notation, mas property accessors

	// map[tipoChave]tipoValor
	user := map[string]string{
		"name":     "hiago",
		"lastName": "klabunde",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	// map aninhado
	teacher := map[string]map[string]string{
		"name": {
			"firstName": "John",
			"lastName":  "Doe",
		},
	}

	fmt.Println(teacher["name"]["firstName"])

	// remove entry do map
	delete(teacher, "name")
	fmt.Println(teacher)
}
