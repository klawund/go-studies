package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "aisjdasijd"
		variavel4 string = "asdsasd"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "v5", "v6"

	fmt.Println(variavel5, variavel6)

	// tudo acima funciona para constantes

	const constante string = "const"
	constante2 := "abc"

	fmt.Println(constante, constante2)

	// atribuição múltipla

	variavel5, variavel6 = variavel6, variavel5
}
