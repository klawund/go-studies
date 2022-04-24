package main

import "fmt"

func main() {
	fmt.Println("estruturas de controle")

	numero := 10
	// a utilização de parênteses só é necessária ao haver precedências diferentes, por padrão não é necessário
	// encapsular a condição
	if numero > 15 {
		fmt.Println("maior que 15")
	} else {
		fmt.Println("menor que 15")
	}

	// if init - permite a declaração de uma variável temporária no if
	if numero2 := 15; numero2 > 0 {
		fmt.Println(numero2) // é possível pois utiliza o escopo interno do if
		fmt.Println("maior que 0")
	}

	// fmt.Println(numero2) - não é possível pois fica restrito ao escopo do if
}
