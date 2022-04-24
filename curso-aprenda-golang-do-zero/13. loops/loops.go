package main

import (
	"fmt"
)

func main() {
	// o go só possui o for, o que muda é o tipo de condição passada
	i := 0

	// while
	for i < 10 {
		i++
		fmt.Println(i)
	}

	fmt.Println("-------------")

	// for padrão
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	// for in range - combinação do for in e for of do js
	// a primeira variavel sempre é o índice e a segunda sempre o valor
	names := [...]string{"hiago", "gabriel", "klabunde"}

	for key, value := range names {
		fmt.Println(key, value)
	}

	// usar o blank quando só quiser um dos valores
	for _, value := range names {
		fmt.Println(value)
	}

	// não é necessário utilizar um array/slice, qualquer iterável funciona
	for key, value := range "word" {
		fmt.Println(key, value)         // como o value sempre será um char, o valor printado é o int8/rune do char na ascii
		fmt.Println(key, string(value)) // dessa forma converte e traz a letra
	}

	user := map[string]string{
		"nome":      "hiago",
		"sobrenome": "klabunde",
	}

	for key, value := range user {
		fmt.Println(key + ": " + value)
	}

	// não é possível iterar struct

	// loop infinito - mesmo que while (true)
	for {
		fmt.Println("aa")
	}
}
