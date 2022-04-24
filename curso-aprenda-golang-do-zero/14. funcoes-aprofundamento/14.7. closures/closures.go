package main

import "fmt"

func closure() func() {
	texto := "dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova() // referencia a var "texto" da closure onde foi fabricada, não a "texto" da main
}
